package orchestrator

import (
	"log"

	"github.com/jestebanrods/edu-golang/dispatch/internal/domain"
)

type Orchestrator struct {
	builder  *WorkerBuilder
	services map[string]*domain.Service
	workers  map[string]*Worker
}

func (o *Orchestrator) CreateService(service *domain.Service) {
	log.Println("orchestrator: creating new service")

	uuid := service.UUID
	_, ok := o.services[uuid]

	if ok {
		log.Printf("orchestrator: service %s alreasy exists", uuid)
	}

	_, ok = o.workers[service.UUID]
	if !ok {
		wk := o.createWorker(service)
		wk.Start()
	}
}

func (o *Orchestrator) ShopperSearchCompleted(id string, shoppers []*domain.Shopper) {
	log.Printf("orchestrator: shopper search completed %s", id)

	wk, ok := o.workers[id]
	if ok {
		wk.ShopperSearchCompleted(shoppers)
	}
}

func (o *Orchestrator) createWorker(service *domain.Service) *Worker {
	log.Printf("orchestrator: building worker %s", service.UUID)

	wk := o.builder.Build(service)

	o.services[service.UUID] = service
	o.workers[service.UUID] = wk

	go wk.Run()

	log.Printf("orchestrator: worker %s running", service.UUID)

	return wk
}

func NewOrchestrator(builder *WorkerBuilder) *Orchestrator {
	return &Orchestrator{
		builder:  builder,
		services: make(map[string]*domain.Service),
		workers:  make(map[string]*Worker),
	}
}
