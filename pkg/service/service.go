package service

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repository"
)

type Authorization interface {
	CreateUser(user dto.User) (dto.User, error)
	Login(origin dto.UserLogin) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
