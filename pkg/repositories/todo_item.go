package repositories

import (
	"errors"
	"fmt"
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories/models"
	"gorm.io/gorm"
)

type TodoItemDB struct {
	db *gorm.DB
}

func (repo TodoItemDB) checkTodoList(userId int, id int) (int, error) {
	var checkIndex int

	row := repo.db.
		Model(&models.TodoList{}).
		Where("user_id = ?", userId).
		Where("id = ?", id).
		Select("todo_lists.id").
		Row()

	err := row.Scan(&checkIndex)

	if err != nil {
		return 0, errors.New(fmt.Sprintf("No todo_list with id %v", id))
	}

	return checkIndex, nil
}

func (repo TodoItemDB) CreateTodoItem(userId int, todoItem dto.TodoItem) (dto.TodoItem, error) {
	_, err := repo.checkTodoList(userId, todoItem.ListId)
	if err != nil {
		return todoItem, err
	}

	var maxOrder int

	row := repo.db.
		Model(&models.TodoItem{}).
		Joins("JOIN todo_lists ON todo_lists.id = todo_items.list_id").
		Select("MAX(item_order) as max_order").
		Where("list_id = ?", todoItem.ListId).
		Row()

	err = row.Scan(&maxOrder)

	if err != nil {
		maxOrder = 0
	}

	maxOrder += 1

	newTodoItem := models.TodoItem{
		ListId:      todoItem.ListId,
		Title:       todoItem.Title,
		Description: todoItem.Description,
		IsDone:      false,
		DoneUntil:   todoItem.DoneUntil,
		Order:       maxOrder,
	}

	result := repo.db.Create(&newTodoItem)
	return newTodoItem.ToDTO(), result.Error
}

func (repo TodoItemDB) GetTodoItems(userId int, todoListId int) ([]dto.TodoItem, error) {
	_, err := repo.checkTodoList(userId, todoListId)
	if err != nil {
		return nil, err
	}

	var todoItems []models.TodoItem
	var resultTodoItems []dto.TodoItem

	result := repo.db.
		Where("list_id = ?", todoListId).
		Order("item_order").
		Order("id").
		Find(&todoItems)

	if result.Error != nil {
		return nil, result.Error
	}
	for _, todoItem := range todoItems {
		resultTodoItems = append(resultTodoItems, todoItem.ToDTO())
	}

	return resultTodoItems, result.Error
}

func (repo TodoItemDB) getTodoItem(userId int, todoListId int, id int) (*models.TodoItem, error) {
	_, err := repo.checkTodoList(userId, todoListId)
	if err != nil {
		return &models.TodoItem{}, err
	}

	var todoItem *models.TodoItem

	result := repo.db.
		Model(&models.TodoItem{}).
		Where("list_id = ?", todoListId).
		Where("id = ?", id).
		First(&todoItem)

	return todoItem, result.Error
}

func (repo TodoItemDB) GetTodoItem(userId int, todoListId int, id int) (dto.TodoItem, error) {
	todoItem, err := repo.getTodoItem(userId, todoListId, id)
	return todoItem.ToDTO(), err
}

func (repo TodoItemDB) UpdateTodoItem(userId int, id int, todoItem dto.TodoItem) (dto.TodoItem, error) {
	updatedTodoItem, err := repo.getTodoItem(userId, todoItem.ListId, id)
	if err != nil {
		return updatedTodoItem.ToDTO(), err
	}

	updatedTodoItem.Title = todoItem.Title
	updatedTodoItem.Description = todoItem.Description
	updatedTodoItem.IsDone = todoItem.IsDone
	updatedTodoItem.DoneUntil = todoItem.DoneUntil

	result := repo.db.Save(&updatedTodoItem)
	return updatedTodoItem.ToDTO(), result.Error
}

func (repo TodoItemDB) ChangeTodoItemOrder(userId, todoListId int, id int, order int) (dto.TodoItem, error) {
	todoItem, err := repo.getTodoItem(userId, todoListId, id)
	if err != nil {
		return todoItem.ToDTO(), err
	}

	err = repo.db.Transaction(func(trans *gorm.DB) error {
		result := trans.
			Model(&models.TodoItem{}).
			Where("item_order BETWEEN ? AND ?", todoItem.Order, order).
			UpdateColumn(
				"item_order",
				gorm.Expr("item_order - ?", 1),
			)

		if result.Error != nil {
			return result.Error
		}

		todoItem.Order = order

		result = trans.Save(&todoItem)

		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return todoItem.ToDTO(), err
}

func (repo TodoItemDB) RemoveTodoItem(userId int, todoListId int, id int) error {
	todoItem, err := repo.getTodoItem(userId, todoListId, id)
	if err != nil {
		return err
	}

	repo.db.Unscoped().Delete(&todoItem)
	return nil
}

func NewTodoItemDB(db *gorm.DB) *TodoItemDB {
	return &TodoItemDB{db}
}
