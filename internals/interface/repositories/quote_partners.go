package repositories

import "github.com/capitual/valete_ms/internals/models"

type IQuotePartnersRepository interface {
	Add(p models.QuotePartner) (models.QuotePartner, error)
}
