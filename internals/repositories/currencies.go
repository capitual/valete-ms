package repositories

import (
	"github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
)

type CurrencyRepository struct{}

func (s *CurrencyRepository) Add(currency models.Currency) (models.Currency, error) {
	db := infra.GetDatabase()

	err := db.Create(&currency).Error

	if err != nil {
		return models.Currency{}, err
	}

	return currency, nil
}

func (s *CurrencyRepository) List(search string) ([]models.Currency, error) {

	db := infra.GetDatabase()

	var currencies []models.Currency

	err := db.Where(`name ILIKE '%` + search + `%'`).Or(`iso_code ILIKE '%` + search + `%'`).Find(&currencies).Error

	if err != nil {
		return []models.Currency{}, nil
	}

	return currencies, nil
}
