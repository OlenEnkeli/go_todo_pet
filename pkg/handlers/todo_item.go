package handlers

import (
	"fmt"
	"github.com/OlenEnkeli/go_todo_pet/pkg/handlers/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ItemUriParams struct {
	ListId int `uri:"id" binding:"required"`
}

type ItemIdUriParams struct {
	ItemUriParams
	ItemId int `uri:"item_id" binding:"required"`
}

type ItemIdOrderIdUriParams struct {
	ItemIdUriParams
	Order int `uri:"order" binding:"required"`
}

func (h *Handler) createTodoItem(ctx *gin.Context) {
	var input schemas.TodoItemCreateSchema
	var uriParams ItemUriParams
	var result schemas.TodoItemReturnSchema

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(
			ctx,
			http.StatusUnprocessableEntity,
			"Missing uri params id",
		)
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	todoItem, err := h.services.CreateTodoItem(
		ctx.GetInt("userId"),
		uriParams.ListId,
		input.ToDTO(),
	)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	result.FromDTO(todoItem)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) getTodoItems(ctx *gin.Context) {
	var uriParams ItemUriParams
	var result schemas.TodoItemsReturnSchema

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Missing uri params id")
		return
	}

	todoItems, err := h.services.GetTodoItems(
		ctx.GetInt("userId"),
		uriParams.ListId,
	)

	if err != nil {
		RaiseErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	for _, todoItem := range todoItems {
		var todoItemSchema schemas.TodoItemReturnSchema
		todoItemSchema.FromDTO(todoItem)

		result.Items = append(result.Items, todoItemSchema)
	}

	result.Amount = len(result.Items)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) getTodoItem(ctx *gin.Context) {
	var uriParams ItemIdUriParams
	var result schemas.TodoItemReturnSchema

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Missing uri params id / item_id")
		return
	}

	todoItem, err := h.services.GetTodoItem(
		ctx.GetInt("userId"),
		uriParams.ListId,
		uriParams.ItemId,
	)
	if err != nil {
		RaiseErrorResponse(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("No todo_list with id %v", uriParams.ItemId),
		)
		return
	}

	result.FromDTO(todoItem)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) updateTodoItem(ctx *gin.Context) {
	var uriParams ItemIdUriParams
	var input schemas.TodoItemUpdateSchema
	var result schemas.TodoItemReturnSchema

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Missing uri params id / item_id")
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	updatedDTO := input.ToDTO()
	updatedDTO.ListId = uriParams.ListId

	todoList, err := h.services.UpdateTodoItem(
		ctx.GetInt("userId"),
		uriParams.ListId,
		uriParams.ItemId,
		updatedDTO,
	)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result.FromDTO(todoList)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) changeTodoItemOrder(ctx *gin.Context) {
	var uriParams ItemIdOrderIdUriParams
	var result schemas.TodoItemReturnSchema

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Required uri params id / item_id / order")
		return
	}

	todoList, err := h.services.ChangeTodoItemOrder(
		ctx.GetInt("userId"),
		uriParams.ListId,
		uriParams.ItemId,
		uriParams.Order,
	)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result.FromDTO(todoList)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) deleteTodoItem(ctx *gin.Context) {
	var uriParams ItemIdUriParams

	err := ctx.ShouldBindUri(&uriParams)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, "Missing uri params id / item_id")
		return
	}

	if err := h.services.RemoveTodoItem(
		ctx.GetInt("userId"),
		uriParams.ListId,
		uriParams.ItemId,
	); err != nil {
		RaiseErrorResponse(
			ctx,
			http.StatusNotFound,
			fmt.Sprintf("Can`t delete todo list with id %v", uriParams.ListId),
		)
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"success": "ok"})
}
