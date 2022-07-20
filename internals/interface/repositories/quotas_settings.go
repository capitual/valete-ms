package repositories

import "github.com/capitual/valete_ms/internals/models"

type IQuoteSettingRepository interface {
	Add(q models.QuotesSetting) (models.QuotesSetting, error)
	// GetById(id uint) (models.QuotasSetting, error)
}
