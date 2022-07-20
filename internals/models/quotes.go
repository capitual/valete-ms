package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Quote struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	QuoteId   string    `json:"quote_id" gorm:"not null"`
	Trading   string    `json:"trading" validate:"required,min=3,max=20" gorm:"not null"`
	Side      string    `json:"side" gorm:"not null"`
	Size      string    `json:"size" gorm:"not null"`
	Price     string    `json:"price" gorm:"not null"`
	ExpiresAt int       `json:"expires_at" gorm:"not null"`
	ExpiresIn int       `json:"expires_in" gorm:"not null"`
	PartnerId uint      `json:"partner_id" gorm:"index,not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (quote *Quote) Prepare() error {
	err := quote.validate()

	if err != nil {
		return err
	}

	return nil
}

func (quote *Quote) validate() error {
	validate := validator.New()
	return validate.Struct(quote)
}
