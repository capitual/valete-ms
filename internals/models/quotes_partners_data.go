package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuotesPartnersData struct {
	ID             uint          `json:"id" gorm:"primaryKey"`
	Key            string        `json:"key" validate:"required,min=1,max=200" gorm:"not null"`
	Value          string        `json:"value" validate:"required,min=1,max=200" gorm:"not null"`
	QuotePartnerID uint          `json:"quote_partner_id" gorm:"not null"`
	QuotePartner   QuotesPartner `json:"quote_partner"`
	Environment    string        `json:"environment" gorm:"not null"`
	CreatedAt      time.Time     `json:"created_at" gorm:"not null"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"not null"`
}

func (quotesPartnerData *QuotesPartnersData) Prepare() error {
	err := quotesPartnerData.validate()

	if err != nil {
		return err
	}
	return nil
}

func (quotesPartnerData *QuotesPartnersData) validate() error {
	validate := validator.New()
	return validate.Struct(quotesPartnerData)
}
