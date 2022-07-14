package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuoteCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3,max=50"`
	Spread    float64   `json:"spread" gorm:"type:float"`
	Ttls      string    `json:"ttls" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (QuoteCategorie *QuoteCategory) Prepare() error {
	err := QuoteCategorie.validate()

	if err != nil {
		return err
	}

	return nil
}

func (quoteCategory *QuoteCategory) validate() error {
	validate := validator.New()
	return validate.Struct(quoteCategory)
}
