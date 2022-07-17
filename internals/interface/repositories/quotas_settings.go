package repositories

import "github.com/capitual/valete_ms/internals/models"

type IQuotaSettingRepository interface {
	Add(q models.QuotasSetting) (models.QuotasSetting, error)
	// GetById(id uint) (models.QuotasSetting, error)
}
