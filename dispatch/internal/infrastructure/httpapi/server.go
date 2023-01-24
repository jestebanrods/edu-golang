package httpapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/httpapi/handler/health"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/httpapi/handler/services"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/httpapi/handler/shoppers"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/orchestrator"
)

type Server struct {
	orchestrator    *orchestrator.Channel
	httpAddr        string
	engine          *gin.Engine
	shutdownTimeout time.Duration
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())

	s.engine.POST("/services", services.CreateHandler(s.orchestrator))
	s.engine.GET("/services/{id}/available", services.AvailableHandler())
	s.engine.GET("/services/{id}", services.FinderHandler())

	s.engine.GET("/shoppers/{id}/history", shoppers.HistoryHandler())
}

func New(host string, port uint, shutdownTimeout time.Duration, orchestrator *orchestrator.Channel) *Server {
	srv := Server{
		engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
		orchestrator:    orchestrator,
	}

	srv.registerRoutes()

	return &srv
}
