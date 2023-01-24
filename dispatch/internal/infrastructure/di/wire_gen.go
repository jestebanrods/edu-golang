// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"github.com/jestebanrods/edu-golang/dispatch/internal"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/orchestrator"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/selector"
)

// Injectors from wire.go:

func InitializeApp() (*internal.App, error) {
	channel := orchestrator.NewChannel()
	server := NewHttpAPI(channel)
	selectorChannel := selector.NewChannel()
	selectorSelector := selector.NewSelector(channel)
	runner := selector.NewRunner(selectorChannel, selectorSelector)
	workerBuilder := orchestrator.NewWorkerBuilder(selectorChannel)
	orchestratorOrchestrator := orchestrator.NewOrchestrator(workerBuilder)
	orchestratorRunner := orchestrator.NewRunner(channel, orchestratorOrchestrator)
	app := internal.NewApp(server, runner, orchestratorRunner)
	return app, nil
}
