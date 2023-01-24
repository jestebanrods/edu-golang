//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jestebanrods/edu-golang/dispatch/internal"
	"github.com/jestebanrods/edu-golang/dispatch/internal/application"
	"github.com/jestebanrods/edu-golang/dispatch/internal/domain"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/orchestrator"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/persistence"
	"github.com/jestebanrods/edu-golang/dispatch/internal/infrastructure/selector"
)

var superSet = wire.NewSet(
	internal.NewApp,

	// UseCases
	application.NewServicesUseCases,
	application.NewShopperTrackingUseCases,

	orchestrator.NewWorkerBuilder,
	orchestrator.NewChannel,
	orchestrator.NewRunner,
	orchestrator.NewOrchestrator,

	selector.NewChannel,
	selector.NewSelector,
	selector.NewRunner,

	// Infrastructure
	persistence.NewDocumentDBServiceRepository,
	persistence.NewRedisShopperTrackingRepository,

	wire.Bind(new(domain.ShopperTrackingRepositoy), new(*persistence.RedisShopperTrackingRepository)),
	wire.Bind(new(domain.Selector), new(*selector.Channel)),
	wire.Bind(new(domain.Orchestrator), new(*orchestrator.Channel)),

	// Real Providers
	// httpapi.New,

	// Fake Providers
	NewHttpAPI,
)
