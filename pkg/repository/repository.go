package repository

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user dto.User) (dto.User, error)
	GetUserByLogin(login string) (dto.User, error)
}

type TodoList interface {
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
	}
}
