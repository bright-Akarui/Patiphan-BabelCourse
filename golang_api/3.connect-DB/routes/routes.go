package routes

import (
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	articlesGroup := r.Group("api/v1/articles")
	articlesController := controllers.Article{}
	{
		articlesGroup.GET("", articlesController.FindAll)
		articlesGroup.GET("/:id", articlesController.FindOne)
		articlesGroup.POST("/post", articlesController.Create)
	}
}
