package models

import (
	"time"
)

type BaseModel struct {
	Id        int       `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type ModelWithIndex struct {
	Id int `gorm:"primarykey"`
}
