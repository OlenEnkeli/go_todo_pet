package models

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"time"
)

type TodoItem struct {
	BaseModel
	ListId      int
	Title       string `gorm:"not null"`
	Description string
	Order       int
	IsDone      bool `gorm:"default:false"`
	DoneUntil   time.Time
}

func (item *TodoItem) ToDTO() dto.TodoItem {
	return dto.TodoItem{
		Id:          item.Id,
		ListId:      item.ListId,
		Title:       item.Title,
		Description: item.Description,
		IsDone:      item.IsDone,
		DoneUntil:   item.DoneUntil,
		CreatedAt:   item.CreatedAt,
	}
}
