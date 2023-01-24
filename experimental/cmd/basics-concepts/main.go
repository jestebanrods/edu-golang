package main

import (
	"github.com/google/wire"
	basics "github.com/jestebanrods/edu-golang/experimental/internal/basics-concepts"
	"github.com/stretchr/testify/mock"
)

func main() {
	basics.Interfaces()
}

type repositoryMock struct {
	mock.Mock
}

var superSet = wire.NewSet()
