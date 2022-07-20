package repositories

import "github.com/capitual/valete_ms/internals/models"

type ICurrencyRepository interface {
	Add(currency models.Currency) (models.Currency, error)
	List(search string) ([]models.Currency, error)
}
