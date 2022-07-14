package services

import (
	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/interface/repositories"
	"github.com/capitual/valete_ms/internals/models"
)

type QuoteCategoryService struct {
	repository repositories.IQuoteCategoryRepository
}

func (s *QuoteCategoryService) CreateQuoteCategory(data dtos.QuoteCategoryDto) (models.QuoteCategory, error) {
	quote_categories := &models.QuoteCategory{
		Name:   data.Name,
		Spread: data.Spread,
		Ttls:   data.Ttls,
	}
	err := quote_categories.Prepare()

	if err != nil {
		return models.QuoteCategory{}, err
	}
	return s.repository.Add(*quote_categories)
}

func NewQuoteCategoryService(r repositories.IQuoteCategoryRepository) *QuoteCategoryService {
	return &QuoteCategoryService{
		repository: r,
	}
}
