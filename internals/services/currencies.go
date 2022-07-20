package services

import (
	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/interface/repositories"
	"github.com/capitual/valete_ms/internals/models"
)

type CurrencyService struct {
	repository repositories.ICurrencyRepository
}

func (s *CurrencyService) CreateCurrency(dto dtos.CurrencyDto) (models.Currency, error) {
	currency := &models.Currency{
		Name:    dto.Name,
		IsoCode: dto.IsoCode,
		Type:    models.CurrencyType(dto.Type),
	}

	err := currency.Prepare()

	if err != nil {
		return models.Currency{}, err
	}

	return s.repository.Add(*currency)
}

func (s *CurrencyService) ListCurrencies(search string) ([]models.Currency, error) {

	currencies, err := s.repository.List(search)

	if err != nil {
		return []models.Currency{}, err
	}

	return currencies, nil
}

func NewCurrencyService(r repositories.ICurrencyRepository) *CurrencyService {
	return &CurrencyService{
		repository: r,
	}
}
