package migrations

import (
	"github.com/capitual/valete_ms/internals/models"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	db.Table("partners").AutoMigrate(models.Partner{})
	db.Table("quotas_categories").AutoMigrate(models.QuotasCategory{})
	db.Table("quotas_settings").AutoMigrate(models.QuotasSetting{})
	db.Migrator().CreateConstraint(&models.Partner{}, "QuotasSettings")
	db.Migrator().CreateConstraint(&models.Partner{}, "fk_partners_quotas_settings")

}
