package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuotePartner struct {
	ID        uint                 `json:"id" gorm:"primaryKey"`
	Name      string               `json:"name" validate:"required,min=2,max=100" gorm:"not null"`
	Slug      string               `json:"slug" validate:"required,min=2,max=100" gorm:"not null"`
	Active    bool                 `json:"active" gorm:"not null"`
	Data      []QuotePartnersData `gorm:"foreignkey:QuotePartnerID" json:"datas,omitempty"`
	CreatedAt time.Time            `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time            `json:"updated_at" gorm:"not null"`
}

func (quotePartner *QuotePartner) Prepare() error {
	quotePartner.Active = true

	err := quotePartner.validate()

	if err != nil {
		return err
	}
	return nil
}

func (quotePartner *QuotePartner) validate() error {
	validate := validator.New()
	return validate.Struct(quotePartner)
}
