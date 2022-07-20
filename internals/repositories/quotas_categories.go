package repositories

import (
	infra "github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
)

type QuoteCategoriesRepository struct{}

func (s *QuoteCategoriesRepository) Add(q models.QuotesCategory) (models.QuotesCategory, error) {
	db := infra.GetDatabase()

	err := db.Create(&q).Error

	if err != nil {
		return models.QuotesCategory{}, err
	}

	return q, nil
}
