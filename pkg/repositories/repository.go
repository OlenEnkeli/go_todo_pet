package repositories

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user dto.User) (dto.User, error)
	GetUserByLogin(login string) (dto.User, error)
	GetCurrentUser(userId int) (dto.User, error)
}

type TodoList interface {
	GetTodoList(userId int, id int) (dto.TodoList, error)
	GetTodoLists(userId int) ([]dto.TodoList, error)
	CreateTodoList(todoList dto.TodoList) (dto.TodoList, error)
	RemoveTodoList(userId int, id int) error
	UpdateTodoList(id int, todoList dto.TodoList) (dto.TodoList, error)
	ChangeTodoListOrder(userId, id int, order int) (dto.TodoList, error)
	GetTodoListsStatistics(userId int) (dto.TodoListStatistic, map[int]*dto.TodoListStatistic, error)
}

type TodoItem interface {
	CreateTodoItem(userId int, todoItem dto.TodoItem) (dto.TodoItem, error)
	GetTodoItem(userId int, todoListId int, id int) (dto.TodoItem, error)
	GetTodoItems(userId int, todoListId int) ([]dto.TodoItem, error)
	UpdateTodoItem(userId int, id int, todoItem dto.TodoItem) (dto.TodoItem, error)
	ChangeTodoItemOrder(userId, todoListId int, id int, order int) (dto.TodoItem, error)
	RemoveTodoItem(userId int, todoListId int, id int) error
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
		TodoItem:      NewTodoItemDB(db),
	}
}
