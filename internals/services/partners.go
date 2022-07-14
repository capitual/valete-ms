package services

import (
	"github.com/capitual/valete_ms/internals/dtos"
	"github.com/capitual/valete_ms/internals/interface/repositories"
	"github.com/capitual/valete_ms/internals/models"
)

type PartnerService struct {
	repository repositories.IPartnerRepository
}

func (s *PartnerService) CreatePartner(data dtos.PartnerDto) (models.Partner, error) {
	partner := &models.Partner{
		PartnerName: data.PartnerName,
		Country:     data.Country,
		Locale:      data.Locale,
	}

	err := partner.Prepare()

	if err != nil {
		return models.Partner{}, err
	}

	return s.repository.Add(*partner)
}

func NewPartnerService(r repositories.IPartnerRepository) *PartnerService {
	return &PartnerService{
		repository: r,
	}
}
