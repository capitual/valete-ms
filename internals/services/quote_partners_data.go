package services

import (
	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/interface/repositories"
	"github.com/capitual/valete_ms/internals/models"
)

type QuotePartnersDataService struct {
	repository repositories.IquotePartnersDataRepository
}

func (s *QuotePartnersDataService) CreateQuotePartnerData(dto dtos.QuotePartnerDataDto) (models.QuotePartnersData, error) {
	data := &models.QuotePartnersData{
		Key:            dto.Key,
		Value:          dto.Value,
		QuotePartnerID: dto.QuotePartnerID,
		Environment:    dto.Environment,
	}

	err := data.Prepare()

	if err != nil {
		return models.QuotePartnersData{}, err
	}

	return s.repository.Add(*data)
}

func NewQuotePartnersDataService(r repositories.IquotePartnersDataRepository) *QuotePartnersDataService {
	return &QuotePartnersDataService{
		repository: r,
	}
}
