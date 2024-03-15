package handlers

import (
	"github.com/OlenEnkeli/go_todo_pet/pkg/handlers/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTodoList(ctx *gin.Context) {
	var input schemas.TodoListCreateSchema

	if err := ctx.ShouldBindJSON(&input); err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	todoList, err := h.services.CreateTodoList(ctx.GetInt("userId"), input.ToDTO())
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var result schemas.TodoListReturnSchema
	result.FromDTO(todoList)

	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) getTodoLists(ctx *gin.Context) {
	todoLists, err := h.services.GetTodoLists(ctx.GetInt("userId"))

	if err != nil {
		RaiseErrorResponse(ctx, http.StatusNotFound, err.Error())
	}

	var result schemas.TodoListsReturnSchema

	for _, todoList := range todoLists {
		var todoListSchema schemas.TodoListReturnSchema
		todoListSchema.FromDTO(todoList)

		result.Items = append(result.Items, todoListSchema)
	}

	result.Amount = len(result.Items)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) getTodoList(ctx *gin.Context) {

}

func (h *Handler) updateTodoList(ctx *gin.Context) {

}

func (h *Handler) deleteTodoList(ctx *gin.Context) {

}
