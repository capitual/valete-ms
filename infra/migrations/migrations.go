package migrations

import (
	"github.com/capitual/valete_ms/internals/models"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	db.Table("partners").AutoMigrate(models.Partner{})
	db.Table("quote_categories").AutoMigrate(models.QuoteCategory{})
}
