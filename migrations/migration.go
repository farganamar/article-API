package migrations

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	MigrateArticles(db)

	// Additional migrations can be added here in the future

}
