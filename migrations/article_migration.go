package migrations

import (
	"article/models"

	"gorm.io/gorm"
)

func MigrateArticles(db *gorm.DB) error {
	return db.AutoMigrate(&models.Article{})
}
