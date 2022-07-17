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

	new_partner, err := s.repository.Add(*partner)

	if err != nil {
		return models.Partner{}, err
	}

	new_partner.PartnerId = ""
	new_partner.PartnerKey = ""

	return new_partner, nil
}

func (s *PartnerService) GetPartnerById(id int) (models.Partner, error) {
	return s.repository.GetById(id)
}

func (s *PartnerService) RevogatePartner(id int) (bool, error) {
	partner, err := s.GetPartnerById(id)

	if err != nil {
		return false, err
	}

	partner.Active = false
	_, err = s.repository.Update(partner)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *PartnerService) GetAll(filters string, page int, size int) (models.PartnerList, error) {
	partner_list, err := s.repository.GetAll(filters, page, size)

	if err != nil {
		return models.PartnerList{}, err
	}

	return partner_list, nil
}

func NewPartnerService(r repositories.IPartnerRepository) *PartnerService {
	return &PartnerService{
		repository: r,
	}
}
