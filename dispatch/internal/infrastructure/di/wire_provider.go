package di

import (
	"time"

	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/httpapi"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/orchestrator"
)

func NewHttpAPI(orchestrator *orchestrator.Channel) *httpapi.Server {
	return httpapi.New("localhost", 9000, 10*time.Second, orchestrator)
}
