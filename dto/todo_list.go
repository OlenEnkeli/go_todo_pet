package dto

import "time"

type TodoList struct {
	Id          int
	UserId      int
	Title       string
	Description string
	Order       int
	CreatedAt   time.Time
}
