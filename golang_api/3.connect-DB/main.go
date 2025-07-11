package main

import (
	"go-gin/config"
	"go-gin/migrations"
	"go-gin/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		// สำหรับใช้งาน env
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// สำหรับใช้งาน เชื่อม db
	config.InitDB()
	defer config.CloseDB()

	// สำหรับรัน Migration เพื่อสร้างหรือเปลี่ยนแปลงฐานข้อมูล
	migrations.Migrate()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	// allow ให้ใครก็ได้สามารถส่ง Header:Authorization เข้ามาได้
	corsConfig.AddAllowHeaders("Authorization")

	r := gin.Default()
	r.Use(cors.New(corsConfig))
	//"http://127.0.0.1:8080/uploads/articles/6/MainBefore.jpg"
	//r.Static("ภายใต้ path ที่เข้ามา","Folder อะไร")
	// สำหรับใช้งาน การเข้าถึง
	r.Static("/uploads", "./uploads")

	//สำหรับสร้าง folder uploads ไฟล์
	uploadDirs := [...]string{"articles", "users"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}
	routes.Serve(r)

	// สำหรับรัน port เท่าไหร่
	r.Run(":" + os.Getenv("PORT"))
}
