package persistence

import "github.com/jestebanrods/edu-golang/dispatch/internal/domain"

type DocumentDBServiceRepository struct {
}

func (r *DocumentDBServiceRepository) FindByUuid(uuid string) *domain.Service {
	return &domain.Service{
		UUID: uuid,
	}
}

func NewDocumentDBServiceRepository() *DocumentDBServiceRepository {
	return &DocumentDBServiceRepository{}
}
