package shoppers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HistoryHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "find history")
	}
}
