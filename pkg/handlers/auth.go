package handlers

import (
	"github.com/OlenEnkeli/go_todo_pet/dto"
	"github.com/OlenEnkeli/go_todo_pet/pkg/handlers/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input schemas.UserCreateSchema

	if err := ctx.ShouldBindJSON(&input); err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user, err := h.services.CreateUser(input.ToDTO())
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var result schemas.UserReturnSchema
	result.FromDTO(user)
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) login(ctx *gin.Context) {
	var input dto.UserLogin

	if err := ctx.ShouldBindJSON(&input); err != nil {
		RaiseErrorResponse(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	token, err := h.services.Login(input)
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token,
	})
}

func (h *Handler) GetCurrentUser(ctx *gin.Context) {
	user, err := h.services.GetCurrentUser(ctx.GetInt("userId"))
	if err != nil {
		RaiseErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var result schemas.UserReturnSchema
	result.FromDTO(user)
	ctx.JSON(http.StatusOK, result)
}
