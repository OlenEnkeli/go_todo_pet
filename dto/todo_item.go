package dto

import "time"

type TodoItem struct {
	Id          uint
	Title       string
	Description string
	IsDone      bool
	DoneUntil   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type ListToItem struct {
	Id     uint
	ListId int
	ItemId int
}
