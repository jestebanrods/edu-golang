package domain

type Selector interface {
	FindShoppers(service Service)
	StopShoppersFind(id string)
}

type Orchestrator interface {
	ShopperSearchCompleted(id string, shoppers []*Shopper)
	CreateService(service *Service)
}
