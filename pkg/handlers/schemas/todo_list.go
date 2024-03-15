package schemas

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"time"
)

type TodoListBaseSchema struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}

type TodoListCreateSchema struct {
	TodoListBaseSchema
}

func (schema *TodoListCreateSchema) ToDTO() dto.TodoList {
	return dto.TodoList{
		Title:       schema.Title,
		Description: schema.Description,
		Order:       schema.Order,
	}
}

type TodoListReturnSchema struct {
	Id     int `json:"id" binding:"required"`
	UserId int `json:"user_id" binding:"required"`
	TodoListBaseSchema
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

type TodoListsReturnSchema struct {
	Amount int                    `json:"amount"`
	Items  []TodoListReturnSchema `json:"items"`
}

func (schema *TodoListReturnSchema) FromDTO(input dto.TodoList) {
	schema.Id = input.Id
	schema.UserId = input.UserId
	schema.Title = input.Title
	schema.Order = input.Order
	schema.Description = input.Description
	schema.CreatedAt = input.CreatedAt
}
