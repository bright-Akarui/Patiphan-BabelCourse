package main

import (
	"go-gin/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	//"http://127.0.0.1:8080/uploads/articles/6/MainBefore.jpg"
	//r.Static("ภายใต้ path ที่เข้ามา","Folder อะไร")
	r.Static("/uploads", "./uploads")
	uploadDirs := [...]string{"articles", "users"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	routes.Serve(r)
	r.Run(":" + os.Getenv("PORT"))
}
