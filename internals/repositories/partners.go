package repositories

import (
	"fmt"
	"math"

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

	err := db.First(&partner, "id = ?", id).Error

	if err != nil {
		return models.Partner{}, err
	}

	return partner, nil
}

func (s *PartnersRepository) Update(p models.Partner) (models.Partner, error) {
	db := infra.GetDatabase()
	db.Save(&p)
	return p, nil
}

func (s *PartnersRepository) GetAll(filters interface{}, page int, size int) (models.PartnerList, error) {
	db := infra.GetDatabase()

	var p []*models.Partner
	err := db.Limit(size).Offset((page - 1) * size).Find(&p).Error

	fmt.Println(p)
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
		TotalCount: int(math.Ceil(float64(count) / float64(size))),
		TotalPages: ((int(count)) / size) + 1,
	}, nil
}
