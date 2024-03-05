package handler

import (
	"github.com/OlenEnkeli/go_todo_pet/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", h.signUp)
		auth.POST("/login", h.login)
	}

	lists := router.Group("/lists")
	{
		lists.POST("/", h.createList)
		lists.GET("/", h.getLists)
		lists.GET("/:id", h.getList)
		lists.PATCH("/:id", h.updateList)
		lists.DELETE("/:id", h.deleteList)

		items := lists.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getItems)
			items.GET("/:item_id", h.getItem)
			items.PATCH("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
