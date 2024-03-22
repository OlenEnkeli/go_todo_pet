package repositories

import (
	"errors"
	"fmt"
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories/models"
	"gorm.io/gorm"
	"time"
)

type TodoListDB struct {
	db *gorm.DB
}

type RowStatistic struct {
	Id      int
	Counter int
}

func (repo TodoListDB) CreateTodoList(todoList dto.TodoList) (dto.TodoList, error) {
	var maxOrder int

	row := repo.db.
		Model(&models.TodoList{}).
		Select("MAX(list_order) as max_order").
		Row()

	err := row.Scan(&maxOrder)

	if err != nil {
		maxOrder = 1
	} else {
		maxOrder += 1
	}

	newTodoList := models.TodoList{
		UserId:      todoList.UserId,
		Title:       todoList.Title,
		Description: todoList.Description,
		Order:       maxOrder,
	}

	result := repo.db.Create(&newTodoList)
	return newTodoList.ToDTO(), result.Error
}

func (repo TodoListDB) GetTodoLists(userId int) ([]dto.TodoList, error) {
	var todoLists []models.TodoList
	var resultTodoLists []dto.TodoList

	result := repo.db.
		Where("user_id = ?", userId).
		Order("list_order").
		Order("id").
		Find(&todoLists)

	if result.Error != nil {
		return nil, result.Error
	}
	for _, todoList := range todoLists {
		resultTodoLists = append(resultTodoLists, todoList.ToDTO())
	}

	return resultTodoLists, result.Error
}

func (repo TodoListDB) getTodoList(userId int, id int) (*models.TodoList, error) {
	var todoList *models.TodoList

	result := repo.db.
		Model(&models.TodoList{}).
		Where("user_id = ?", userId).
		Where("id = ?", id).
		First(&todoList)

	return todoList, result.Error
}

func (repo TodoListDB) GetTodoList(userId int, id int) (dto.TodoList, error) {
	todoList, err := repo.getTodoList(userId, id)
	return todoList.ToDTO(), err
}

func (repo TodoListDB) statisticRequest(
	userId int,
	where string,
	statistic *dto.TodoListStatistic,
	statistics map[int]*dto.TodoListStatistic,
	targetField string,
) error {
	var rowStatistics []RowStatistic

	query := fmt.Sprintf(`
		SELECT todo_lists.id AS id, COUNT(todo_items.id) AS counter
		FROM todo_items 
		JOIN todo_lists ON todo_lists.id = todo_items.list_id
		WHERE todo_lists.user_id = ?
		AND %s
		GROUP BY todo_lists.id
		ORDER BY todo_lists.id;
	`, where)

	result := repo.db.
		Raw(
			query,
			userId,
		).
		Scan(&rowStatistics)

	if result.Error != nil {
		return result.Error
	}

	for _, rowStatistic := range rowStatistics {
		if _, ok := statistics[rowStatistic.Id]; ok == false {
			statistics[rowStatistic.Id] = &dto.TodoListStatistic{}
		}

		switch targetField {
		case "ItemAmount":
			statistics[rowStatistic.Id].ItemAmount = rowStatistic.Counter
			statistic.ItemAmount += rowStatistic.Counter
		case "DoneItemAmount":
			statistics[rowStatistic.Id].DoneItemAmount = rowStatistic.Counter
			statistic.DoneItemAmount += rowStatistic.Counter
		case "PlannedItemAmount":
			statistics[rowStatistic.Id].PlannedItemAmount = rowStatistic.Counter
			statistic.PlannedItemAmount += rowStatistic.Counter
		case "TodayItemAmount":
			statistics[rowStatistic.Id].TodayItemAmount = rowStatistic.Counter
			statistic.TodayItemAmount += rowStatistic.Counter
		case "ExpiredItemAmount":
			statistics[rowStatistic.Id].ExpiredItemAmount = rowStatistic.Counter
			statistic.ExpiredItemAmount += rowStatistic.Counter
		default:
			return errors.New("wrong target field")
		}
	}

	return nil
}

func (repo TodoListDB) GetTodoListsStatistics(userId int) (dto.TodoListStatistic, map[int]*dto.TodoListStatistic, error) {
	statistics := make(map[int]*dto.TodoListStatistic)
	statistic := dto.TodoListStatistic{
		ItemAmount:        0,
		DoneItemAmount:    0,
		PlannedItemAmount: 0,
		TodayItemAmount:   0,
		ExpiredItemAmount: 0,
	}

	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	yesterday := now.AddDate(0, 0, -1)
	tomorrowStart := time.Date(
		tomorrow.Year(),
		tomorrow.Month(),
		tomorrow.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	yesterdayStart := time.Date(
		yesterday.Year(),
		yesterday.Month(),
		yesterday.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)

	err := repo.statisticRequest(
		userId,
		"true",
		&statistic,
		statistics,
		"ItemAmount",
	)
	if err != nil {
		return dto.TodoListStatistic{}, nil, err
	}

	err = repo.statisticRequest(
		userId,
		"todo_items.is_done IS true",
		&statistic,
		statistics,
		"DoneItemAmount",
	)
	if err != nil {
		return dto.TodoListStatistic{}, nil, err
	}

	err = repo.statisticRequest(
		userId,
		"todo_items.is_done IS false",
		&statistic,
		statistics,
		"PlannedItemAmount",
	)
	if err != nil {
		return dto.TodoListStatistic{}, nil, err
	}

	err = repo.statisticRequest(
		userId,
		fmt.Sprintf(
			"todo_items.is_done IS false AND todo_items.done_until < '%s' AND todo_items.done_until > '%s'",
			toSQLDateTime(tomorrowStart),
			toSQLDateTime(yesterdayStart),
		),
		&statistic,
		statistics,
		"TodayItemAmount",
	)
	if err != nil {
		return dto.TodoListStatistic{}, nil, err
	}

	err = repo.statisticRequest(
		userId,
		fmt.Sprintf(
			"todo_items.is_done IS false AND todo_items.done_until < '%s'",
			toSQLDateTime(now),
		),
		&statistic,
		statistics,
		"ExpiredItemAmount",
	)
	if err != nil {
		return dto.TodoListStatistic{}, nil, err
	}

	return statistic, statistics, nil
}

func (repo TodoListDB) UpdateTodoList(id int, todoList dto.TodoList) (dto.TodoList, error) {
	updatedTodoList, err := repo.getTodoList(todoList.UserId, id)
	if err != nil {
		return todoList, err
	}

	updatedTodoList.Title = todoList.Title
	updatedTodoList.Description = todoList.Description

	result := repo.db.Save(&updatedTodoList)
	return updatedTodoList.ToDTO(), result.Error
}

func (repo TodoListDB) ChangeTodoListOrder(userId, id int, order int) (dto.TodoList, error) {
	todoList, err := repo.getTodoList(userId, id)
	if err != nil {
		return todoList.ToDTO(), err
	}

	err = repo.db.Transaction(func(trans *gorm.DB) error {
		if todoList.Order == order {
			return nil
		}

		var updateFrom int
		var updateUntil int
		var updateExpr string

		if todoList.Order < order {
			updateFrom = todoList.Order
			updateUntil = order
			updateExpr = "list_order - ?"
		} else {
			updateFrom = order
			updateUntil = todoList.Order
			updateExpr = "list_order + ?"
		}

		result := trans.
			Model(&models.TodoList{}).
			Where("list_order BETWEEN ? AND ?", updateFrom, updateUntil).
			UpdateColumn(
				"list_order",
				gorm.Expr(updateExpr, 1),
			)

		if result.Error != nil {
			return result.Error
		}

		todoList.Order = order

		result = trans.Save(&todoList)

		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return todoList.ToDTO(), err
}

func (repo TodoListDB) RemoveTodoList(userId int, id int) error {
	todoList, err := repo.getTodoList(userId, id)
	if err != nil {
		return err
	}

	repo.db.Unscoped().Delete(&todoList)
	return nil
}

func NewTodoListDB(db *gorm.DB) *TodoListDB {
	return &TodoListDB{db: db}
}
