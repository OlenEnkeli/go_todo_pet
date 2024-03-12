package services

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories"
)

type Authorization interface {
	CreateUser(user dto.User) (dto.User, error)
	Login(origin dto.UserLogin) (string, error)
	ParseToken(token string) (uint, error)
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

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
