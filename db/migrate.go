package db

import (
	"hydra/db/extensions"
	"hydra/db/models"
	"hydra/db/types"

	"gorm.io/gorm"
)


func RunMigrations(db *gorm.DB) {
	extensions.MigrateExtensions(db)
	types.MigrateTypes(db)
	models.MigrateModels(db)
}