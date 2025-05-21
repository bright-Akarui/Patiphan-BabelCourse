package migrations

import (
	"go-gin/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1747844730CreateArticlesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1747844730",
		Migrate: func(tx *gorm.DB) error {
			//ตรวจสอบว่ามี Column ID_Category อยู่ไหม ถ้าไม่มีก็จะถูกเพิ่มเข้ามาสำหรับ Auto migration
			return tx.AutoMigrate(&models.Article{}) // จะ return error มาให้เลย
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropColumn(&models.Article{}, "ID_Category")
		},
	}
}
