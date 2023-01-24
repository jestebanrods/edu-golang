package selector

import (
	"log"

	"github.com/jestebanrods/edu-golang/dispatch/internal/domain"
)

type FindStartEvent struct {
	service domain.Service
}

type FindStopEvent struct {
	ID string
}

type Channel struct {
	channel chan interface{}
}

func (c *Channel) FindShoppers(service domain.Service) {
	log.Println("selector-channel: find shoppers")
	c.channel <- FindStartEvent{
		service: service,
	}
}

func (c *Channel) StopShoppersFind(id string) {
	log.Println("selector-channel: stop finding shoppers")
	c.channel <- FindStopEvent{
		ID: id,
	}
}

func NewChannel() *Channel {
	return &Channel{
		channel: make(chan interface{}, 100),
	}
}
