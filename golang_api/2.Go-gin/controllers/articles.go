package controllers

import (
	"go-gin/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var articles []models.Article = []models.Article{
	{ID: 1, Title: "text#2", Body: "Body#2"},
	{ID: 2, Title: "text#3", Body: "Body#3"},
	{ID: 3, Title: "text#4", Body: "Body#4"},
	{ID: 4, Title: "text#1", Body: "Body#1"},
	{ID: 5, Title: "text#5", Body: "Body#5"},
}

type createArticle struct {
	Title string                `form:"title" binding:"required"`
	Body  string                `form:"body" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}
type Article struct {
}

func (a *Article) FindAll(ctx *gin.Context) {
	result := articles
	// ctx.Query("limit") หมายความว่า หา key ที่ชื่อ limit และเอาค่ามา ยกตัวอย่างเช่น article?limit=3
	// limit = 3
	limit := ctx.Query("limit")
	if limit != "" {
		n, _ := strconv.Atoi(limit)
		result = result[:n]
	}
	ctx.JSON(http.StatusOK, gin.H{"articles": result})
}

func (a *Article) FindOne(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for _, v := range articles {
		if v.ID == uint(id) {
			ctx.JSON(http.StatusOK, gin.H{"articles": v})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, "Notfound")
}
func (a *Article) Create(ctx *gin.Context) {
	var form createArticle
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ar := models.Article{
		ID:    uint(len(articles) + 1),
		Title: form.Title,
		Body:  form.Body,
	}
	//Get file
	file, _ := ctx.FormFile("image")
	//Create Path
	path := "uploads/articles/" + strconv.Itoa(int(ar.ID))
	os.MkdirAll(path, 0775)
	//Upload file
	filename := path + "/" + file.Filename
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		return
	}
	// Attach file to article
	ar.Image = os.Getenv("HOST") + "/" + filename
	articles = append(articles, ar)
	ctx.JSON(http.StatusCreated, gin.H{"articles": articles})
}
