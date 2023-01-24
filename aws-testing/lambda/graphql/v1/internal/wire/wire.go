// +build wireinject

package wire

import (
	"github.com/jestebanrods/edu-golang/aws-testing/lambda/graphql/v1/internal"

	"github.com/google/wire"
)

func Initialize() (*internal.Handler, error) {
	wire.Build(stdSet)
	return &internal.Handler{}, nil
}
