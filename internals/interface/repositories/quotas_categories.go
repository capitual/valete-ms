package repositories

import "github.com/capitual/valete_ms/internals/models"

type IQuotaCategoryRepository interface {
	Add(q models.QuotasCategory) (models.QuotasCategory, error)
}
