package migrations

import (
	"go-gin/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1747737106CreateArticlesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1747737106",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Article{}) // จะ return error มาให้เลย
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("articles") // จะ return error มาให้เลย
		},
	}
}
