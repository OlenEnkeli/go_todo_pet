package handlers

import (
	"fmt"
	"github.com/OlenEnkeli/go_todo_pet/pkg/handlers/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListIdUriParams struct {
	Id int `uri:"id" binding:"required"`
}

type ListIdOrderIdUriParams struct {
	ListIdUriParams
	Order int `uri:"order" binding:"required"`
}

func (h *Handler) createTodoList(ctx *gin.Context) {
	var input schemas.TodoListCreateSchema
	var result schemas.TodoListReturnSchema

	if err := ctx.ShouldBindJSON(&input); err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	todoList, err := h.services.CreateTodoList(ctx.GetInt("userId"), input.ToDTO())
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result.FromDTO(todoList)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) getTodoLists(ctx *gin.Context) {
	todoLists, err := h.services.GetTodoLists(ctx.GetInt("userId"))
	var result schemas.TodoListsReturnSchema

	if err != nil {
		RaiseErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	for _, todoList := range todoLists {
		var todoListSchema schemas.TodoListReturnSchema
		todoListSchema.FromDTO(todoList)

		result.Items = append(result.Items, todoListSchema)
	}

	result.Amount = len(result.Items)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) getTodoList(ctx *gin.Context) {
	var uriParams ListIdUriParams
	var result schemas.TodoListReturnSchema

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Missing uri params id")
		return
	}

	todoList, err := h.services.GetTodoList(ctx.GetInt("userId"), uriParams.Id)
	if err != nil {
		RaiseErrorResponse(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("No todo_list with id %v", uriParams.Id),
		)
		return
	}

	result.FromDTO(todoList)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) updateTodoList(ctx *gin.Context) {
	var uriParams ListIdUriParams
	var input schemas.TodoListUpdateSchema
	var result schemas.TodoListReturnSchema

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Missing uri params id")
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	updatedDTO := input.ToDTO()
	updatedDTO.UserId = ctx.GetInt("userId")

	todoList, err := h.services.UpdateTodoList(uriParams.Id, updatedDTO)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result.FromDTO(todoList)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) changeTodoListOrder(ctx *gin.Context) {
	var uriParams ListIdOrderIdUriParams
	var result schemas.TodoListReturnSchema

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Required uri params id / order")
		return
	}

	todoList, err := h.services.ChangeTodoListOrder(
		ctx.GetInt("userId"),
		uriParams.Id,
		uriParams.Order,
	)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result.FromDTO(todoList)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) deleteTodoList(ctx *gin.Context) {
	var uriParams ListIdUriParams

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Missing uri params id")
		return
	}

	if err := h.services.RemoveTodoList(ctx.GetInt("userId"), uriParams.Id); err != nil {
		RaiseErrorResponse(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("Can`t delete todo list with id %v", uriParams.Id),
		)
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"success": "ok"})
}
