package models

import (
	"time"

	"github.com/capitual/valete_ms/internals/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Partner struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PartnerName string    `json:"partner_name" validate:"required,min=3,max=50"`
	PartnerKey  string    `json:"partner_key" validate:"required"`
	PartnerId   string    `json:"partner_id" validate:"required"`
	Country     string    `json:"country" validate:"required,min=3,max=50"`
	Locale      string    `json:"locale" validate:"required,min=2,max=5"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (partner *Partner) Prepare() error {
	partner.PartnerId = uuid.New().String()
	partner.PartnerKey = utils.RandomString(50)
	partner.Active = true

	err := partner.validate()
	if err != nil {
		return err
	}
	return nil
}

func (partner *Partner) validate() error {
	validate := validator.New()
	return validate.Struct(partner)
}
