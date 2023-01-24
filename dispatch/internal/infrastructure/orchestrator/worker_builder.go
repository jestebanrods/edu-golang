package orchestrator

import (
	"github.com/jestebanrods/edu-golang/dispatch/internal/domain"
)

type WorkerBuilder struct {
	selector domain.Selector
}

func NewWorkerBuilder(selector domain.Selector) *WorkerBuilder {
	return &WorkerBuilder{
		selector: selector,
	}
}

func (ob *WorkerBuilder) Build(service *domain.Service) *Worker {
	return &Worker{
		service:  service,
		events:   make(chan interface{}, 100),
		selector: ob.selector,
	}
}
