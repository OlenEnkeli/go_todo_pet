package dto

import "time"

type TodoItem struct {
	Id          int
	ListId      int
	Title       string
	Description string
	Order       int
	IsDone      bool
	DoneUntil   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
