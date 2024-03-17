package repositories

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories/models"
	"gorm.io/gorm"
)

type TodoListDB struct {
	db *gorm.DB
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
		result := trans.
			Model(&models.TodoList{}).
			Where("list_order BETWEEN ? AND ?", todoList.Order, order).
			UpdateColumn(
				"list_order",
				gorm.Expr("list_order - ?", 1),
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
