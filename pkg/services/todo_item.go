package services

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories"
)

type TodoItemService struct {
	repo repositories.TodoItem
}

func (srv *TodoItemService) CreateTodoItem(userId int, todoListId int, todoItem dto.TodoItem) (dto.TodoItem, error) {
	todoItem.ListId = todoListId
	return srv.repo.CreateTodoItem(userId, todoItem)
}

func (srv *TodoItemService) GetTodoItem(userId int, todoListId int, id int) (dto.TodoItem, error) {
	return srv.repo.GetTodoItem(userId, todoListId, id)
}

func (srv *TodoItemService) GetTodoItems(userId int, todoListId int) ([]dto.TodoItem, error) {
	return srv.repo.GetTodoItems(userId, todoListId)
}

func (srv *TodoItemService) RemoveTodoItem(userId int, todoListId int, id int) error {
	return srv.repo.RemoveTodoItem(userId, todoListId, id)
}

func (srv *TodoItemService) UpdateTodoItem(userId int, todoListId int, id int, todoItem dto.TodoItem) (dto.TodoItem, error) {
	todoItem.ListId = todoListId
	return srv.repo.UpdateTodoItem(userId, id, todoItem)
}

func (srv *TodoItemService) ChangeTodoItemOrder(userId, todoListId int, id int, order int) (dto.TodoItem, error) {
	return srv.repo.ChangeTodoItemOrder(userId, todoListId, id, order)
}

func NewTodoItemService(repo repositories.TodoItem) *TodoItemService {
	return &TodoItemService{repo}
}
