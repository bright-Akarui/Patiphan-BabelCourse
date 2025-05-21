package migrations

import (
	"go-gin/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1747841916CreateArticlesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1747841917",
		Migrate: func(tx *gorm.DB) error {
			//ตรวจสอบว่ามี tables ชื่อ articles หรือยังไ ถ้ายังไม่มีจะทำการสร้าง table ขึ้นมา
			return tx.AutoMigrate(&models.Category{}) // จะ return error มาให้เลย
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("categories") // จะ return error มาให้เลย
		},
	}
}
