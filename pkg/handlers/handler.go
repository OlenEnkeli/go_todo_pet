package handlers

import (
	"github.com/OlenEnkeli/go_todo_pet/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/users", h.signUp)
	currentUser := router.Group("/users/current", h.authNeeded)
	{
		currentUser.GET("", h.GetCurrentUser)
	}
	router.POST("/login", h.login)

	lists := router.Group("/lists", h.authNeeded)
	{
		lists.POST("/", h.createTodoList)
		lists.GET("/", h.getTodoLists)
		lists.GET("/:id", h.getTodoList)
		lists.PATCH("/:id", h.updateTodoList)
		lists.DELETE("/:id", h.deleteTodoList)
		lists.PATCH("/:id/order/:order", h.changeTodoListOrder)

		items := lists.Group(":id/items")
		{
			items.POST("/", h.createTodoItem)
			items.GET("/", h.getTodoItems)
			items.GET("/:item_id", h.getTodoItem)
			items.PATCH("/:item_id", h.updateTodoItem)
			items.DELETE("/:item_id", h.deleteTodoItem)
			items.PATCH("/:item_id/order/:order", h.changeTodoItemOrder)
		}
	}

	return router
}
