package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FinderHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "find by id")
	}
}
