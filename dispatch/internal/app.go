package internal

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/httpapi"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/orchestrator"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/selector"
)

type App struct {
	http         *httpapi.Server
	selector     *selector.Runner
	orchestrator *orchestrator.Runner
}

func (s *App) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	go s.selector.Run()
	go s.orchestrator.Run()
	go s.http.Run(ctx)

	s.Stop()

	return nil
}

func (s *App) Stop() {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	signal := <-sign

	log.Println("syscall", signal.String())
	log.Println("salida controlada")
}

func NewApp(http *httpapi.Server, selector *selector.Runner, orchestrator *orchestrator.Runner) *App {
	return &App{
		http:         http,
		selector:     selector,
		orchestrator: orchestrator,
	}
}
