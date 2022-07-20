package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuotesPartner struct {
	ID        uint                 `json:"id" gorm:"primaryKey"`
	Name      string               `json:"name" validate:"required,min=2,max=100" gorm:"not null"`
	Slug      string               `json:"slug" validate:"required,min=2,max=100" gorm:"not null"`
	Active    bool                 `json:"active" gorm:"not null"`
	Data      []QuotesPartnersData `gorm:"foreignkey:QuotePartnerID" json:"datas,omitempty"`
	CreatedAt time.Time            `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time            `json:"updated_at" gorm:"not null"`
}

func (quotesPartner *QuotesPartner) Prepare() error {
	quotesPartner.Active = true

	err := quotesPartner.validate()

	if err != nil {
		return err
	}
	return nil
}

func (quotesPartner *QuotesPartner) validate() error {
	validate := validator.New()
	return validate.Struct(quotesPartner)
}
