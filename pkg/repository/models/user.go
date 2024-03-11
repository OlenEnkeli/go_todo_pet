package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey"`
	Login    string `gorm:"index:user_login_index,unique"`
	Name     string
	Password string
}
