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

func (s *PartnersRepository) GetById(id int) (models.Partner, error) {
	db := infra.GetDatabase()

	var partner models.Partner

	err := db.Select(`id, partner_name, country, locale, 
	active, created_at, updated_at`).First(&partner, "id = ?", id).Error

	if err != nil {
		return models.Partner{}, err
	}

	return partner, nil
}

func (s *PartnersRepository) Update(p models.Partner) (models.Partner, error) {
	db := infra.GetDatabase()
	err := db.Save(&p).Error

	if err != nil {
		return models.Partner{}, err
	}

	return p, nil
}

func (s *PartnersRepository) GetAll(filter string, page int, size int) (models.PartnerList, error) {
	db := infra.GetDatabase()

	var p []*models.Partner
	err := db.Limit(size).Offset((page - 1) * size).Select(`id, partner_name, country, locale, 
	active, created_at, updated_at`).Where(`partner_name ILIKE '%` + filter + `%'`).Find(&p).Error

	if err != nil {
		return models.PartnerList{}, err
	}

	var count int64

	err = db.Model(&models.Partner{}).Count(&count).Error

	if err != nil {
		return models.PartnerList{}, err
	}

	return models.PartnerList{
		Body:       p,
		TotalCount: int(count),
		TotalPages: ((int(count)) / size) + 1,
	}, nil
}
