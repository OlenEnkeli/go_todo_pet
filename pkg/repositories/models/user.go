package models

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
)

type User struct {
	BaseModel
	Login     string     `gorm:"index:user_login_index,unique; not null"`
	Name      string     `gorm:"not null"`
	Password  string     `gorm:"not null"`
	TodoLists []TodoList `gorm:"foreignKey:UserId"`
}

func (user *User) ToDTO() dto.User {
	result := dto.User{
		Id:       user.Id,
		Login:    user.Login,
		Name:     user.Name,
		Password: user.Password,
	}

	return result
}
