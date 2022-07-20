package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuotePartnersData struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Key            string    `json:"key" validate:"required,min=1,max=200" gorm:"not null"`
	Value          string    `json:"value" validate:"required,min=1,max=200" gorm:"not null"`
	QuotePartnerID uint      `json:"quote_partner_id" gorm:"not null"`
	Environment    string    `json:"environment" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"not null"`
}

func (quotePartnerData *QuotePartnersData) Prepare() error {
	err := quotePartnerData.validate()

	if err != nil {
		return err
	}
	return nil
}

func (quotePartnerData *QuotePartnersData) validate() error {
	validate := validator.New()
	return validate.Struct(quotePartnerData)
}
