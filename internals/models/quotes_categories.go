package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type QuotesCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3,max=50" gorm:"not null"`
	SpreadAsk string    `json:"spread_ask" validate:"required" gorm:"not null"`
	SpreadBid string    `json:"spread_bid" validate:"required" gorm:"not null"`
	Ttls      int       `json:"ttls" validate:"required" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (quoteCategory *QuotesCategory) Prepare() error {
	err := quoteCategory.validate()

	if err != nil {
		return err
	}

	return nil
}

func (quoteCategory *QuotesCategory) validate() error {
	validate := validator.New()
	return validate.Struct(quoteCategory)
}
