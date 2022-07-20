package models

import (
	"time"

	"github.com/capitual/valete_ms/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Partner struct {
	ID          uint            `json:"id" gorm:"primaryKey"`
	PartnerName string          `json:"partner_name" validate:"required,min=3,max=50" gorm:"not null"`
	PartnerKey  string          `json:"partner_key,omitempty" validate:"required" gorm:"not null"`
	PartnerId   string          `json:"partner_id,omitempty" validate:"required" gorm:"not null"`
	Country     string          `json:"country" validate:"required,min=3,max=50" gorm:"not null"`
	Locale      string          `json:"locale" validate:"required,min=2,max=5" gorm:"not null"`
	Active      bool            `json:"active" gorm:"not null"`
	Settings    []QuotesSetting `json:"quotes_settings,omitempty" gorm:"foreignKey:PartnerId"`
	Quotes      []Quote         `json:"quotes,omitempty" gorm:"foreignKey:PartnerId"`
	CreatedAt   time.Time       `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"not null"`
}

func (partner *Partner) Prepare() error {
	partner.PartnerId = uuid.New().String()
	partner.PartnerKey = utils.RandomString(50)
	partner.Active = true

	err := partner.validate()
	if err != nil {
		return err
	}
	return nil
}

func (partner *Partner) validate() error {
	validate := validator.New()
	return validate.Struct(partner)
}

type PartnerList struct {
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
	// HasMore    bool       `json:"has_more"`
	Body []*Partner `json:"data"`
}
