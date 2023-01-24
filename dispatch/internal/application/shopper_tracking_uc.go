package application

import "github.com/jestebanrods/edu-golang/dispatch/internal/domain"

type ShopperTrackingUseCases struct {
	repository domain.ShopperTrackingRepositoy
}

func (uc *ShopperTrackingUseCases) FindAvailable() []*domain.ShopperTracking {
	return uc.repository.FindByStatus(domain.ShopperStateAvailable)
}

func NewShopperTrackingUseCases(repository domain.ShopperTrackingRepositoy) *ShopperTrackingUseCases {
	return &ShopperTrackingUseCases{
		repository: repository,
	}
}
