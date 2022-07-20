package repositories

import "github.com/capitual/valete_ms/internals/models"

type IquotePartnersDataRepository interface {
	Add(p models.QuotePartnersData) (models.QuotePartnersData, error)
}
