package repositories

import (
	"github.com/OlenEnkeli/go_todo_pet/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetDBConnection(config *configs.ConfigStruct) *gorm.DB {
	logrus.Infof("DB: %s", config.GetPostgresDSN())

	db, err := gorm.Open(
		postgres.Open(config.GetPostgresDSN()),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalf("DB Connection Error: %s", err)
	}

	return db
}
