package repositories

import (
	"github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
)

type QuotePartnersDataRepository struct{}

func (s *QuotePartnersDataRepository) Add(p models.QuotePartnersData) (models.QuotePartnersData, error) {
	db := infra.GetDatabase()

	err := db.Create(&p).Error

	if err != nil {
		return models.QuotePartnersData{}, err
	}

	return p, nil
}
