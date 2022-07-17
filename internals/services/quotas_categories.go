package services

import (
	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/interface/repositories"
	"github.com/capitual/valete_ms/internals/models"
)

type QuotaCategoryService struct {
	repository repositories.IQuotaCategoryRepository
}

func (s *QuotaCategoryService) CreateQuotaCategory(data dtos.QuotaCategoryDto) (models.QuotasCategory, error) {
	quote_categories := &models.QuotasCategory{
		Name:   data.Name,
		Spread: data.Spread,
		Ttls:   data.Ttls,
	}
	err := quote_categories.Prepare()

	if err != nil {
		return models.QuotasCategory{}, err
	}
	return s.repository.Add(*quote_categories)
}

func NewQuotaCategoryService(r repositories.IQuotaCategoryRepository) *QuotaCategoryService {
	return &QuotaCategoryService{
		repository: r,
	}
}
