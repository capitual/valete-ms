package repositories

import "github.com/capitual/valete_ms/internals/models"

type IQuoteCategoryRepository interface {
	Add(q models.QuotesCategory) (models.QuotesCategory, error)
}
