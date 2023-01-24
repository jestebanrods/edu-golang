package selector

import (
	"log"

	"github.com/jestebanrods/edu-golang/dispatch/internal/domain"
)

type Selector struct {
	orchestrator domain.Orchestrator
}

func (s *Selector) FindShoppers(service domain.Service) []*domain.Shopper {
	log.Println("selector: find shoppers")
	return nil
}

func (s *Selector) StopShoppersFind(id string) {
	log.Println("selector: stop shoppers find")
}

func (s *Selector) NotifySearch(id string, shoppers []*domain.Shopper) {
	log.Println("selector: notify search")
	s.orchestrator.ShopperSearchCompleted(id, shoppers)
}

func NewSelector(orchestrator domain.Orchestrator) *Selector {
	return &Selector{
		orchestrator: orchestrator,
	}
}
