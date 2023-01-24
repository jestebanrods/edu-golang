//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jestebanrods/edu-golang/dispatch/internal"
)

func InitializeApp() (*internal.App, error) {
	wire.Build(superSet)
	return &internal.App{}, nil
}
