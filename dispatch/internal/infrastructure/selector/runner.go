package selector

import (
	"log"
)

type Runner struct {
	ch       <-chan interface{}
	selector *Selector
}

func (r *Runner) Run() {
	log.Println("runner selector up")

	for event := range r.ch {
		switch evt := event.(type) {
		case FindStartEvent:
			log.Println("selector-runner: find start event")
			shoppers := r.selector.FindShoppers(evt.service)
			r.selector.NotifySearch(evt.service.UUID, shoppers)
		case FindStopEvent:
			log.Println("selector-runner: find stop event")
			r.selector.StopShoppersFind(evt.ID)
		}
	}

	log.Println("runner selector down")
}

func NewRunner(ch *Channel, selector *Selector) *Runner {
	return &Runner{
		ch:       ch.channel,
		selector: selector,
	}
}
