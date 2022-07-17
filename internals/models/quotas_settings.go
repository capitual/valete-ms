package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuotasSetting struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Currency  string    `json:"currency" validate:"required,min=3,max=30"`
	ExpiresIn int       `json:"expires_in" validate:"required"`
	PartnerId uint      `json:"partner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (quotaSetting *QuotasSetting) Prepare() error {
	err := quotaSetting.validate()
	if err != nil {
		return err
	}

	return nil
}

func (quotaCategory *QuotasSetting) validate() error {
	validate := validator.New()
	return validate.Struct(quotaCategory)
}
