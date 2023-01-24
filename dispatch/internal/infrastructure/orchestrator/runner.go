package orchestrator

import (
	"log"
)

type Runner struct {
	ch           <-chan interface{}
	orchestrator *Orchestrator
}

func (r *Runner) Run() {
	log.Println("runner orchestrator up")

	for event := range r.ch {
		switch evt := event.(type) {
		case CreateServiceEvent:
			log.Println("orchestrator-runner: create service event")
			r.orchestrator.CreateService(evt.service)
		case SearchingCompleteEvent:
			log.Println("orchestrator-runner: searching complete event")
			r.orchestrator.ShopperSearchCompleted(evt.ID, evt.shoppers)
		}
	}

	log.Println("runner orchestrator down")
}

func NewRunner(ch *Channel, orchestrator *Orchestrator) *Runner {
	return &Runner{
		ch:           ch.channel,
		orchestrator: orchestrator,
	}
}
