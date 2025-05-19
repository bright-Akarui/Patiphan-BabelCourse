package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() {
	var err error
	var logLevel logger.LogLevel
	if gin.Mode() == gin.DebugMode {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}

	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_CONNECTION")), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("failed to get sql.DB:", err)
		return
	}
	sqlDB.Close()
}
