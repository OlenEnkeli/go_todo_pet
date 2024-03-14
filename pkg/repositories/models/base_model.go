package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int            `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ModelWithIndex struct {
	Id int `gorm:"primarykey"`
}
