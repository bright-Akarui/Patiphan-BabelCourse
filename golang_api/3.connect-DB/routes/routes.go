package routes

import (
	"go-gin/config"
	"go-gin/controllers"
	"go-gin/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	// authorize := middleware.Authorize()
	// ประกาศ middleware ก่อน
	authMiddleware := middleware.Authenticate()
	if err := authMiddleware.MiddlewareInit(); err != nil {
		log.Fatal("JWT Middleware Init Error:", err)
	}

	authenticate := middleware.Authenticate().MiddlewareFunc()

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
		categoriesGroup.POST("/post", categoriesController.CreatCategory)
	}
	userGroup := r.Group("api/v1/user")
	userController := controllers.Auth{DB: db}
	{
		userGroup.POST("/create", userController.Signup)
		userGroup.POST("/sign-in", authMiddleware.LoginHandler)
		userGroup.GET("/profile", authenticate, userController.GetProfile)
	}
}
