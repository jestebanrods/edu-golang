package domain

type ShopperTrackingRepositoy interface {
	FindById(id string) *ShopperTracking
	FindByStatus(status uint16) []*ShopperTracking
}

type ShopperTracking struct {
	ID     string
	Status int
}
