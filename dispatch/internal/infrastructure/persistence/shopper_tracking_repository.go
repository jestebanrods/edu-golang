package persistence

import "github.com/jestebanrods/edu-golang/dispatch/internal/domain"

type RedisShopperTrackingRepository struct {
}

func (r *RedisShopperTrackingRepository) FindById(id string) *domain.ShopperTracking {
	return nil
}

func (r *RedisShopperTrackingRepository) FindByStatus(status uint16) []*domain.ShopperTracking {
	return nil
}

func NewRedisShopperTrackingRepository() *RedisShopperTrackingRepository {
	return &RedisShopperTrackingRepository{}
}
