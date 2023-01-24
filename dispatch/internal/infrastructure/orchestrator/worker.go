package orchestrator

import (
	"log"

	"github.com/jestebanrods/edu-golang/dispatch/internal/domain"
)

type FindStartDispatcherEvent struct {
}

type SearchCompletedDispatcherEvent struct {
	shoppers []*domain.Shopper
}

type Worker struct {
	service  *domain.Service
	events   chan interface{}
	selector domain.Selector
}

func (o *Worker) Start() {
	o.events <- FindStartDispatcherEvent{}
}

func (o *Worker) ShopperSearchCompleted(shoppers []*domain.Shopper) {
	o.events <- SearchCompletedDispatcherEvent{
		shoppers: shoppers,
	}
}

func (o *Worker) Run() {
	log.Printf("dispatcher: running for %s", o.service.UUID)

	for event := range o.events {
		// switch evt := event.(type) {
		switch event.(type) {
		case FindStartDispatcherEvent:
			log.Println("dispatcher: init find shoppers")
			o.selector.FindShoppers(*o.service)
		case SearchCompletedDispatcherEvent:
			log.Println("dispatcher: shoppers shearch complete")
			// TODO publish
			log.Println("dispatcher: todo publish services")
		}
	}
}
