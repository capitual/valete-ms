package repositories

import (
	infra "github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
)

type QuotaCategoriesRepository struct{}

func (s *QuotaCategoriesRepository) Add(q models.QuotasCategory) (models.QuotasCategory, error) {
	db := infra.GetDatabase()

	err := db.Create(&q).Error

	if err != nil {
		return models.QuotasCategory{}, err
	}

	return q, nil
}
