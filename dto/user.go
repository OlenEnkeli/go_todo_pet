package dto

import "time"

type User struct {
	Id        int
	Name      string
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserLogin struct {
	Login    string
	Password string
}
