package schemas

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"time"
)

type TodoListBaseSchema struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type TodoListCreateSchema struct {
	TodoListBaseSchema
}

func (schema *TodoListCreateSchema) ToDTO() dto.TodoList {
	return dto.TodoList{
		Title:       schema.Title,
		Description: schema.Description,
	}
}

type TodoListReturnSchema struct {
	Id     int `json:"id" binding:"required"`
	UserId int `json:"user_id" binding:"required"`
	TodoListBaseSchema
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

func (schema *TodoListReturnSchema) FromDTO(input dto.TodoList) {
	schema.Id = input.Id
	schema.UserId = input.UserId
	schema.Title = input.Title
	schema.Order = input.Order
	schema.Description = input.Description
	schema.CreatedAt = input.CreatedAt
}

type TodoListUpdateSchema struct {
	TodoListBaseSchema
}

func (schema *TodoListUpdateSchema) ToDTO() dto.TodoList {
	return dto.TodoList{
		Title:       schema.Title,
		Description: schema.Description,
	}
}

type TodoListsReturnSchema struct {
	Amount int                    `json:"amount"`
	Items  []TodoListReturnSchema `json:"items"`
}
