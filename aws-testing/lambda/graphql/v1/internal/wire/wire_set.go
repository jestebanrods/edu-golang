// +build wireinject

package wire

import (
	"github.com/jestebanrods/edu-golang/aws-testing/lambda/graphql/v1/internal"

	"github.com/google/wire"
)

var stdSet = wire.NewSet(
	internal.NewHandler,
	internal.NewPersonRepositoy,
	internal.NewQueries,
	SchemaFileProvider,
	SchemaProvider,
)
