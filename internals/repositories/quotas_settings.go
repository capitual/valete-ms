package repositories

import (
	infra "github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
)

type QuoteSettingRepository struct{}

func (s *QuoteSettingRepository) Add(q models.QuotesSetting) (models.QuotesSetting, error) {
	db := infra.GetDatabase()

	err := db.Create(&q).Error

	if err != nil {
		return models.QuotesSetting{}, err
	}

	return q, nil
}
