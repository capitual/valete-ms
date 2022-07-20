package services

import (
	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/interface/repositories"
	"github.com/capitual/valete_ms/internals/models"
)

type QuoteSettingService struct {
	repository repositories.IQuoteSettingRepository
}

func (s *QuoteSettingService) CreateQuoteSetting(dto dtos.QuoteSettingDto) (models.QuotesSetting, error) {
	quoteSetting := &models.QuotesSetting{
		PartnerId:  dto.PartnerId,
		ExpiresIn:  dto.ExpiresIn,
		CurrencyId: dto.Currency,
	}

	err := quoteSetting.Prepare()
	if err != nil {
		return models.QuotesSetting{}, err
	}

	return s.repository.Add(*quoteSetting)
}

func NewQuoteSettingService(r repositories.IQuoteSettingRepository) *QuoteSettingService {
	return &QuoteSettingService{
		repository: r,
	}
}
