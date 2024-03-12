package main

import (
	todo "github.com/OlenEnkeli/go_todo_pet"
	"github.com/OlenEnkeli/go_todo_pet/configs"
	"github.com/OlenEnkeli/go_todo_pet/pkg/handlers"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repositories"
	"github.com/OlenEnkeli/go_todo_pet/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"slices"
)

func main() {
	configs.InitConfig()

	if slices.Contains(
		[]string{"dev", "test"},
		configs.Config.Common.Mode,
	) {
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	db := repositories.GetDBConnection(&configs.Config)

	repos := repositories.NewRepository(db)
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services)

	switch configs.Config.Common.Mode {
	case "dev":
		gin.SetMode("debug")
	case "prod":
		gin.SetMode("release")
	default:
		gin.SetMode("debug")
	}

	logrus.Infof(
		"Starting server at http://%s:%s [%s mode]",
		configs.Config.Server.Host,
		configs.Config.Server.Port,
		configs.Config.Common.Mode,
	)

	srv := new(todo.Server)
	if err := srv.Run(
		handlers.InitRoutes(),
		configs.Config.Server.Host,
		configs.Config.Server.Port,
	); err != nil {
		logrus.Fatalf("Fatal error: %s", err)
	}
}
