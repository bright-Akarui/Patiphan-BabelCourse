package routes

import (
	"go-gin/config"
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	articlesGroup := r.Group("api/v1/articles")
	articlesController := controllers.Article{DB: db}
	{
		articlesGroup.GET("", articlesController.FindAll)
		articlesGroup.GET("/:id", articlesController.ArticleFindById)
		articlesGroup.POST("/post", articlesController.Create)
	}
}
