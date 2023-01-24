package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AvailableHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "return services")
	}
}
