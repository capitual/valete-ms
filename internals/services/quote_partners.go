package services

import (
	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/interface/repositories"
	"github.com/capitual/valete_ms/internals/models"
)

type QuotePartnersService struct {
	repository repositories.IQuotePartnersRepository
}

func (s *QuotePartnersService) CreateQuotePartner(dto dtos.QuotePartnerDto) (models.QuotePartner, error) {
	data := &models.QuotePartner{
		Name: dto.Name,
		Slug: dto.Slug,
	}

	err := data.Prepare()

	if err != nil {
		return models.QuotePartner{}, err
	}

	return s.repository.Add(*data)
}

func NewQuotePartnersService(r repositories.IQuotePartnersRepository) *QuotePartnersService {
	return &QuotePartnersService{
		repository: r,
	}
}
