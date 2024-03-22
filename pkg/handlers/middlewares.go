package handlers

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) authNeeded(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		RaiseErrorResponse(ctx, 401, "Auth headers is not set")
		ctx.Abort()
		return
	}

	headersParts := strings.Split(header, " ")
	if len(headersParts) != 2 {
		RaiseErrorResponse(ctx, 401, "Wrong auth header format")
		ctx.Abort()
		return
	}

	userId, err := h.services.ParseToken(headersParts[1])
	if err != nil {
		RaiseErrorResponse(ctx, 401, err.Error())
		ctx.Abort()
		return
	}

	ctx.Set("userId", userId)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
