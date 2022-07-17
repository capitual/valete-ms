package repositories

import (
	"github.com/capitual/valete_ms/internals/models"
)

type IPartnerRepository interface {
	Add(p models.Partner) (models.Partner, error)
	GetById(id int) (models.Partner, error)
	Update(p models.Partner) (models.Partner, error)
	GetAll(filters string, page int, size int) (models.PartnerList, error)
}
