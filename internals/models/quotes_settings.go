package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuotesSetting struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CurrencyId uint      `json:"currency_id" gorm:"not null"`
	ExpiresIn  int       `json:"expires_in" validate:"required" gorm:"not null"`
	PartnerId  uint      `json:"partner_id" gorm:"index,not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"not null"`
}

func (quoteCategory *QuotesSetting) Prepare() error {
	err := quoteCategory.validate()
	if err != nil {
		return err
	}

	return nil
}

func (quoteCategory *QuotesSetting) validate() error {
	validate := validator.New()
	return validate.Struct(quoteCategory)
}
