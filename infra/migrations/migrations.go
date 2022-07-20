package migrations

import (
	"github.com/capitual/valete_ms/internals/models"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	db.Table("partners").AutoMigrate(&models.Partner{})
	db.Table("quotes_categories").AutoMigrate(&models.QuotesCategory{})
	db.Table("currencies").AutoMigrate(&models.Currency{})
	db.Table("quotes_settings").AutoMigrate(&models.QuotesSetting{})
	db.Table("quotes_partners").AutoMigrate(&models.QuotesPartner{})
	db.Table("quotes").AutoMigrate(&models.Quote{})

	db.AutoMigrate(&models.QuotesPartner{}, &models.QuotesPartnersData{})
	db.Migrator().CreateConstraint(&models.Partner{}, "Settings")
	db.Migrator().CreateConstraint(&models.Partner{}, "fk_partners_quotes_settings")
	db.Migrator().CreateConstraint(&models.Currency{}, "Settings")
	db.Migrator().CreateConstraint(&models.Currency{}, "fk_currencies_quotes_settings")
	db.Migrator().CreateConstraint(&models.Partner{}, "Quotes")
	db.Migrator().CreateConstraint(&models.Partner{}, "fk_partners_quotes")
	db.Migrator().CreateConstraint(&models.QuotesPartner{}, "Datas")
	db.Migrator().CreateConstraint(&models.QuotesPartner{}, "fk_quotes_partners_data")
}
