package main

import (
	todo "github.com/OlenEnkeli/go_todo_pet"
	"github.com/OlenEnkeli/go_todo_pet/pkg/handler"
	"github.com/OlenEnkeli/go_todo_pet/pkg/repository"
	"github.com/OlenEnkeli/go_todo_pet/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Fatal error: %d", err)
	}
}
