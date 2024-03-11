package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RaiseErrorResponse(
	ctx *gin.Context,
	statusCode int,
	description string,
) {
	logrus.Warningf(
		"[%s] %s return %d: %s",
		ctx.Request.Method,
		ctx.Request.URL,
		statusCode,
		description,
	)

	ctx.JSON(
		statusCode,
		gin.H{"description:": description},
	)
}
