package models

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"time"
)

type TodoItem struct {
	BaseModel
	ListId      int    `gorm:"index:item_list_id_index"`
	Title       string `gorm:"not null"`
	Description string
	Order       int  `gorm:"column:item_order; index:item_order_index"`
	IsDone      bool `gorm:"default:false"`
	DoneUntil   *time.Time
}

func (item *TodoItem) ToDTO() dto.TodoItem {
	return dto.TodoItem{
		Id:          item.Id,
		ListId:      item.ListId,
		Title:       item.Title,
		Description: item.Description,
		IsDone:      item.IsDone,
		Order:       item.Order,
		DoneUntil:   item.DoneUntil,
		CreatedAt:   item.CreatedAt,
	}
}
