package repositories

import (
	"github.com/OlenEnkeli/go_todo_pet/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func GetDBConnection(config *configs.ConfigStruct) *gorm.DB {
	logrus.Infof("DB: %s", config.GetPostgresDSN())

	var gormConfig gorm.Config

	if config.Common.Mode == "dev" {
		gormConfig = gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	db, err := gorm.Open(
		postgres.Open(config.GetPostgresDSN()),
		&gormConfig,
	)

	if err != nil {
		log.Fatalf("DB Connection Error: %s", err)
	}

	return db
}
