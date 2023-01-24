package wire

import (
	"io/ioutil"

	"github.com/jestebanrods/edu-golang/aws-testing/lambda/graphql/v1/internal"

	graphql "github.com/graph-gophers/graphql-go"
)

type GraphQLSchema string

func SchemaFileProvider() GraphQLSchema {
	dat, err := ioutil.ReadFile("./bin/schema.graphql")
	if err != nil {
		panic(err.Error())
	}

	return GraphQLSchema(string(dat))
}

func SchemaProvider(schema GraphQLSchema, queries *internal.Queries) *graphql.Schema {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	return graphql.MustParseSchema(string(schema), queries, opts...)
}
