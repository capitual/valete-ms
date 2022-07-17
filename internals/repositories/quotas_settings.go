package repositories

import (
	infra "github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
)

type QuotaSettingRepository struct{}

func (s *QuotaSettingRepository) Add(q models.QuotasSetting) (models.QuotasSetting, error) {
	db := infra.GetDatabase()

	err := db.Create(&q).Error

	if err != nil {
		return models.QuotasSetting{}, err
	}

	return q, nil
}
