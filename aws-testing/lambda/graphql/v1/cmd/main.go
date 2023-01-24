package main

import (
	"log"

	"github.com/jestebanrods/edu-golang/aws-testing/lambda/graphql/v1/internal/wire"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler, err := wire.Initialize()
	if err != nil {
		log.Fatalf("fatal err %s", err.Error())
	}

	lambda.Start(handler.Handle)
}
