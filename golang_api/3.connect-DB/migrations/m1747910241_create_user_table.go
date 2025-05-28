package migrations

import (
	"go-gin/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1747910241CreateUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1747910241",
		Migrate: func(tx *gorm.DB) error {
			//ตรวจสอบว่ามี tables ชื่อ User หรือยังไ ถ้ายังไม่มีจะทำการสร้าง table ขึ้นมา
			return tx.AutoMigrate(&models.User{}) // จะ return error มาให้เลย
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users") // จะ return error มาให้เลย
		},
	}
}
