package repositories

import (
	"github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
)

type QuotePartnerRepository struct{}

func (s *QuotePartnerRepository) Add(p models.QuotePartner) (models.QuotePartner, error) {
	db := infra.GetDatabase()

	err := db.Create(&p).Error

	if err != nil {
		return models.QuotePartner{}, nil
	}

	return p, nil
}
