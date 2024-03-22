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

type TodoListStatisticSchema struct {
	ItemAmount        int `json:"task_amount" binding:"required"`
	DoneItemAmount    int `json:"done_task_amount" binding:"required"`
	PlannedItemAmount int `json:"planned_task_amount" binding:"required"`
	TodayItemAmount   int `json:"today_task_amount" binding:"required"`
	ExpiredItemAmount int `json:"expired_task_amount" binding:"required"`
}

type TodoListStatisticReturnSchema struct {
	Statistic      TodoListStatisticSchema         `json:"statistic" binding:"required"`
	ListStatistics map[int]TodoListStatisticSchema `json:"list_statistics" binding:"required"`
}

func (schema *TodoListStatisticReturnSchema) FromDTOs(
	statistic dto.TodoListStatistic,
	statistics map[int]*dto.TodoListStatistic,
) {
	schema.ListStatistics = make(map[int]TodoListStatisticSchema)

	schema.Statistic = TodoListStatisticSchema{
		ItemAmount:        statistic.ItemAmount,
		DoneItemAmount:    statistic.DoneItemAmount,
		PlannedItemAmount: statistic.PlannedItemAmount,
		TodayItemAmount:   statistic.TodayItemAmount,
		ExpiredItemAmount: statistic.ExpiredItemAmount,
	}

	for key, currentStatistic := range statistics {
		schema.ListStatistics[key] = TodoListStatisticSchema{
			ItemAmount:        currentStatistic.ItemAmount,
			DoneItemAmount:    currentStatistic.DoneItemAmount,
			PlannedItemAmount: currentStatistic.PlannedItemAmount,
			TodayItemAmount:   currentStatistic.TodayItemAmount,
			ExpiredItemAmount: currentStatistic.ExpiredItemAmount,
		}
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
