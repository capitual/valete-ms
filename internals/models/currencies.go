package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CurrencyType string

const (
	fiat   CurrencyType = "fiat"
	crypto CurrencyType = "crypto"
	metal  CurrencyType = "metal"
)

type Currency struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name" validate:"required,min=3,max=50" gorm:"not null"`
	IsoCode   string          `json:"iso_code" validate:"required,min=3,max=20" gorm:"not null"`
	Type      CurrencyType    `json:"type" gorm:"not null"`
	Active    bool            `json:"active" gorm:"not null"`
	Settings  []QuotesSetting `json:"quotes_settings,omitempty" gorm:"foreignKey:CurrencyId"`
	CreatedAt time.Time       `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"not null"`
}

func (currency *Currency) Prepare() error {
	currency.Active = true

	err := currency.validate()

	if err != nil {
		return err
	}

	return nil
}

func (currency *Currency) validate() error {
	validate := validator.New()
	return validate.Struct(currency)
}
