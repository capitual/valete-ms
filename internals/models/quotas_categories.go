package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuotasCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3,max=50"`
	Spread    string    `json:"spread"`
	Ttls      int       `json:"ttls" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (quotaCategory *QuotasCategory) Prepare() error {
	err := quotaCategory.validate()

	if err != nil {
		return err
	}

	return nil
}

func (quotaCategory *QuotasCategory) validate() error {
	validate := validator.New()
	return validate.Struct(quotaCategory)
}
