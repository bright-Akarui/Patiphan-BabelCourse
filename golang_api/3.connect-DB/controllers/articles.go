package controllers

import (
	"go-gin/models"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type createArticle struct {
	Title   string                `form:"title" binding:"required"`
	Body    string                `form:"body" binding:"required"`
	Excerpt string                `form:"excerpt" binding:"required"`
	Image   *multipart.FileHeader `form:"image" binding:"required"`
}

type articleSuccess struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
	Body    string `json:"body"`
	Image   string `json:"image"`
}

type Article struct {
	DB *gorm.DB
}

func (a *Article) FindAll(ctx *gin.Context) {
	var articles []models.Article

	if err := a.DB.Find(&articles).Error; err != nil {
		// a.DB.Find(&articles) = SELECT * FROM articles;
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found article"})
		return
	}
	// res := []articleSuccess{}
	var res []articleSuccess
	copier.Copy(&res, &articles)
	ctx.JSON(http.StatusOK, gin.H{"articles": res})
}

func (a *Article) Create(ctx *gin.Context) {
	// ประกาศตัวแปรที่เก็บ struct ของ createArticle
	var form createArticle
	// ดัก error ถ้า ctx.ShouldBind(&form) ไม่เท่ากับ null ให้แสดง key เป็น error และแสดงค่าที่ error ออกไป err.Error()
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// ถ้าไม่ error ให้ทำการดึงข้อมูลมาเก็บใน struck ของ article
	var article models.Article
	// copier.Copy(เอาข้อมูล stuck จาก param2 มาใส่ , copy struck ที่ต้องการไปใส่ใน param1)
	copier.Copy(&article, &form)

	// เชื่อมต่อ DB และส่งข้อมูลไปใน DB
	if err := a.DB.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	errImage := a.setArticleImage(ctx, &article)
	if errImage != nil {
		log.Fatal("error insert image")
		return
	}
	articleSuccess := articleSuccess{}
	copier.Copy(&articleSuccess, &article)
	ctx.JSON(http.StatusCreated, gin.H{"articles": articleSuccess})
}
func (a *Article) setArticleImage(ctx *gin.Context, article *models.Article) error {
	// get image file
	file, err := ctx.FormFile("image")
	// เช็คว่ามี error หรือ file ได้ถูกส่งมามาไหม
	if err != nil || file == nil {
		return err
	}

	// กรณีมีไฟล์อยู่แล้วต้องการอัพเดท
	if article.Image != "" {
		//แล้วค่าเป็นดังนี้ http://127.0.0.1:8080/uploads/articles/<ID>/image.jpg
		//1 ทำให้หรือแค่นี้ /uploads/articles/<ID>/image.jpg
		// strings.Replace(string ต้นฉบับที่ต้องการ replace , อักษรที่ต้องการเปลี่ยน , เปลี่ยนเป็นอะไร , จำทำการ replace กี่ครั้ง)
		article.Image = strings.Replace(article.Image, os.Getenv("HOST"), "", 1)
		//2 กำหนดใหม่ <WD หรือ path ปัจจุบันที่รันอยู่>/uploads/articles/<ID>/image.jpg
		// get <WD>
		pwd, _ := os.Getwd()
		//3 Remove <WD หรือ path ปัจจุบันที่รันอยู่>/uploads/articles/<ID>/image.jpg
		os.Remove(pwd + article.Image)
	}
	//Create Path
	path := "uploads/articles/" + strconv.Itoa(int(article.ID))
	os.MkdirAll(path, 0775)
	//Upload file
	filename := path + "/" + file.Filename
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		return err
	}
	// Attach file to article
	article.Image = os.Getenv("HOST") + "/" + filename
	// update article.image
	a.DB.Save(article)
	return nil
}
func (a *Article) ArticleFindById(ctx *gin.Context) {
	var article models.Article
	id := ctx.Param("id")
	if err := a.DB.First(&article, id).Error; err != nil {
		// SELECT * FROM articles WHERE id = ค่าของ id;
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "not found id"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"article": article})
}
