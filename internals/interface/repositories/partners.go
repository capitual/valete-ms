package repositories

import "github.com/capitual/valete_ms/internals/models"

type IPartnerRepository interface {
	Add(p models.Partner) (models.Partner, error)
}
