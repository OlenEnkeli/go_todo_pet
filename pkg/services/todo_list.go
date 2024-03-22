package services

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories"
)

type TodoListService struct {
	repo repositories.TodoList
}

func (srv *TodoListService) CreateTodoList(userId int, todoList dto.TodoList) (dto.TodoList, error) {
	todoList.UserId = userId
	return srv.repo.CreateTodoList(todoList)
}

func (srv *TodoListService) GetTodoLists(userId int) ([]dto.TodoList, error) {
	return srv.repo.GetTodoLists(userId)
}

func (srv *TodoListService) GetTodoList(userId int, id int) (dto.TodoList, error) {
	return srv.repo.GetTodoList(userId, id)
}

func (srv *TodoListService) UpdateTodoList(id int, todoList dto.TodoList) (dto.TodoList, error) {
	return srv.repo.UpdateTodoList(id, todoList)
}

func (srv *TodoListService) RemoveTodoList(userId int, id int) error {
	return srv.repo.RemoveTodoList(userId, id)
}

func (srv *TodoListService) ChangeTodoListOrder(userId int, id int, order int) (dto.TodoList, error) {
	return srv.repo.ChangeTodoListOrder(userId, id, order)
}

func (srv *TodoListService) GetTodoListsStatistics(userId int) (dto.TodoListStatistic, map[int]*dto.TodoListStatistic, error) {
	return srv.repo.GetTodoListsStatistics(userId)
}

func NewTodoListService(repo repositories.TodoList) *TodoListService {
	return &TodoListService{repo}
}
