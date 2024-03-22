package dto

import "time"

type TodoListStatistic struct {
	ItemAmount        int
	DoneItemAmount    int
	PlannedItemAmount int
	TodayItemAmount   int
	ExpiredItemAmount int
}

type TodoList struct {
	Id          int
	UserId      int
	Title       string
	Description string
	Order       int
	CreatedAt   time.Time
}
