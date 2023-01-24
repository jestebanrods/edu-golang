package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jestebanrods/edu-golang/dispatch/internal/domain"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/orchestrator"
)

type createRequest struct {
	ID string `json:"id" binding:"required"`
}

func CreateHandler(orchestrator *orchestrator.Channel) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		service := domain.Service{
			UUID: req.ID,
		}

		orchestrator.CreateService(&service)

		ctx.Status(http.StatusCreated)
	}
}
