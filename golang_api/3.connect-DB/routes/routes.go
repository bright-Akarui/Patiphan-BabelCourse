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
		articlesGroup.PATCH("/:id", articlesController.UpdateArticle)
		articlesGroup.DELETE("/:id", articlesController.DeleteArticle)
		articlesGroup.POST("/post", articlesController.Create)
	}
	categoriesGroup := r.Group("api/v1/categories")
	categoriesController := controllers.Category{DB: db}
	{
		categoriesGroup.GET("", categoriesController.FindAllCategories)
		// categoriesGroup.GET("/:id", categoriesController.ArticleFindById)
		// categoriesGroup.PATCH("/:id", categoriesController.UpdateArticle)
		// categoriesGroup.DELETE("/:id", categoriesController.DeleteArticle)
		categoriesGroup.POST("/post", categoriesController.CreatCategory)
	}
}
