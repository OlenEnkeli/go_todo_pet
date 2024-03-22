package services

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories"
)

type Authorization interface {
	CreateUser(user dto.User) (dto.User, error)
	Login(origin dto.UserLogin) (string, error)
	ParseToken(token string) (int, error)
	GetCurrentUser(userId int) (dto.User, error)
}

type TodoList interface {
	CreateTodoList(userId int, todoList dto.TodoList) (dto.TodoList, error)
	GetTodoList(userId int, id int) (dto.TodoList, error)
	GetTodoLists(userId int) ([]dto.TodoList, error)
	RemoveTodoList(userId int, id int) error
	UpdateTodoList(id int, todoList dto.TodoList) (dto.TodoList, error)
	ChangeTodoListOrder(userId, id int, order int) (dto.TodoList, error)
	GetTodoListsStatistics(userId int) (dto.TodoListStatistic, map[int]*dto.TodoListStatistic, error)
}

type TodoItem interface {
	CreateTodoItem(userId int, todoListId int, todoItem dto.TodoItem) (dto.TodoItem, error)
	GetTodoItem(userId int, todoListId int, id int) (dto.TodoItem, error)
	GetTodoItems(userId int, todoListId int) ([]dto.TodoItem, error)
	RemoveTodoItem(userId int, todoListId int, id int) error
	UpdateTodoItem(userId int, todoListId int, id int, todoItem dto.TodoItem) (dto.TodoItem, error)
	ChangeTodoItemOrder(userId, todoListId int, id int, order int) (dto.TodoItem, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		TodoList:      NewTodoListService(repos),
		TodoItem:      NewTodoItemService(repos),
	}
}
