package models

import "github.com/OlenEnkeli/go_todo_pet/dto"

type TodoList struct {
	BaseModel
	UserId      int
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (list *TodoList) ToDTO() dto.TodoList {
	return dto.TodoList{
		Id:          list.Id,
		Title:       list.Title,
		Description: list.Description,
	}
}
