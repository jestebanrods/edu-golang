package domain

const (
	ServiceStateCreated uint16 = 200
)

type ServiceRepositoy interface {
	FindByUuid(uuid string) *Service
}

type Service struct {
	UUID        string `bson:"uuid,omitempty"`
	Status      int    `bson:"status,omitempty"`
	Reference   string `bson:"reference,omitempty"`
	CountryCode string `bson:"country_code,omitempty"`
}
