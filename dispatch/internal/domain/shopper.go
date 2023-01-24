package domain

const (
	ShopperStateAvailable uint16 = 200
)

type ShopperRepositoy interface {
	FindById(id string) *Shopper
}

type Shopper struct {
	ID     string
	Status int
}
