package services

import (
	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/interface/repositories"
	"github.com/capitual/valete_ms/internals/models"
)

type QuotaSettingService struct {
	repository repositories.IQuotaSettingRepository
}

func (s *QuotaSettingService) CreateQuotaSetting(dto dtos.QuotaSettingDto) (models.QuotasSetting, error) {
	quotaSetting := &models.QuotasSetting{
		PartnerId: dto.PartnerId,
		ExpiresIn: dto.ExpiresIn,
		Currency:  dto.Currency,
	}

	err := quotaSetting.Prepare()
	if err != nil {
		return models.QuotasSetting{}, err
	}

	return s.repository.Add(*quotaSetting)
}

func NewQuotaSettingService(r repositories.IQuotaSettingRepository) *QuotaSettingService {
	return &QuotaSettingService{
		repository: r,
	}
}
