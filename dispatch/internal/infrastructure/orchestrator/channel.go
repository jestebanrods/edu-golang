package orchestrator

import (
	"log"

	"github.com/jestebanrods/edu-golang/dispatch/internal/domain"
)

type CreateServiceEvent struct {
	service *domain.Service
}

type SearchingCompleteEvent struct {
	ID       string
	shoppers []*domain.Shopper
}

type Channel struct {
	channel chan interface{}
}

func (c *Channel) CreateService(service *domain.Service) {
	log.Println("orchestrator-channel: create service event")
	c.channel <- CreateServiceEvent{
		service: service,
	}
}

func (c *Channel) ShopperSearchCompleted(id string, shoppers []*domain.Shopper) {
	log.Println("orchestrator-channel: finish shopper search complete event")
	c.channel <- SearchingCompleteEvent{
		ID:       id,
		shoppers: shoppers,
	}
}

func NewChannel() *Channel {
	return &Channel{
		channel: make(chan interface{}, 100),
	}
}
