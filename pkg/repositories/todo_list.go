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
	newTodoList := models.TodoList{
		UserId:      todoList.UserId,
		Title:       todoList.Title,
		Description: todoList.Description,
		ListOrder:   todoList.Order,
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

func (repo TodoListDB) GetTodoList(userId int, id int) (dto.TodoList, error) {
	var todoList *models.TodoList

	err := repo.db.
		Model(&models.TodoList{}).
		Where("user_id = ?", userId).
		Where("id = ?", id).
		First(&todoList)

	println(todoList)

	return todoList.ToDTO(), err.Error
}

func (repo TodoListDB) RemoveTodoList(id int) error {
	//TODO implement me
	panic("implement me")
}

func (repo TodoListDB) UpdateTodoList(id int, todoList dto.TodoList) (dto.TodoList, error) {
	//TODO implement me
	panic("implement me")
}

func NewTodoListDB(db *gorm.DB) *TodoListDB {
	return &TodoListDB{db: db}
}
