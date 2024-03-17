package schemas

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"time"
)

type TodoItemBaseSchema struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type TodoItemCreateSchema struct {
	TodoListBaseSchema
}

func (schema *TodoItemCreateSchema) ToDTO() dto.TodoItem {
	return dto.TodoItem{
		Title:       schema.Title,
		Description: schema.Description,
	}
}

type TodoItemReturnSchema struct {
	Id     int `json:"id" binding:"required"`
	ListId int `json:"list_id" binding:"required"`
	TodoItemBaseSchema
	Order     int        `json:"order" binding:"requited"`
	CreatedAt time.Time  `json:"created_at" binding:"required"`
	IsDone    bool       `json:"is_done" binding:"required"`
	DoneUntil *time.Time `json:"done_until" binding:"required"`
}

func (schema *TodoItemReturnSchema) FromDTO(input dto.TodoItem) {
	schema.Id = input.Id
	schema.ListId = input.ListId
	schema.Title = input.Title
	schema.Order = input.Order
	schema.Description = input.Description
	schema.CreatedAt = input.CreatedAt
	schema.IsDone = input.IsDone
	schema.DoneUntil = input.DoneUntil
}

type TodoItemUpdateSchema struct {
	TodoItemBaseSchema
	IsDone    bool       `json:"is_done" binding:"required"`
	DoneUntil *time.Time `json:"done_until"`
}

func (schema *TodoItemUpdateSchema) ToDTO() dto.TodoItem {
	return dto.TodoItem{
		Title:       schema.Title,
		Description: schema.Description,
		IsDone:      schema.IsDone,
		DoneUntil:   schema.DoneUntil,
	}

}

type TodoItemsReturnSchema struct {
	Amount int                    `json:"amount"`
	Items  []TodoItemReturnSchema `json:"items"`
}
