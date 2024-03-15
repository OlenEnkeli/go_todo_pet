package repositories

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user dto.User) (dto.User, error)
	GetUserByLogin(login string) (dto.User, error)
}

type TodoList interface {
	GetTodoList(id int) (dto.TodoList, error)
	GetTodoLists(userId int) ([]dto.TodoList, error)
	CreateTodoList(todoList dto.TodoList) (dto.TodoList, error)
	RemoveTodoList(id int) error
	UpdateTodoList(id int, todoList dto.TodoList) (dto.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDB(db),
		TodoList:      NewTodoListDB(db),
	}
}
