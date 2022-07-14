package repositories

import (
	infra "github.com/capitual/valete_ms/infra"
	"github.com/capitual/valete_ms/internals/models"
)

type PartnersRepository struct{}

func (s *PartnersRepository) Add(p models.Partner) (models.Partner, error) {
	db := infra.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return models.Partner{}, err
	}

	return p, nil
}
