package controllers

import (
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type creatCategory struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc" binding:"required"`
}

type categoriesSuccess struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Category struct {
	DB *gorm.DB
}

func (a *Category) FindAllCategories(ctx *gin.Context) {
	var categories []models.Category

	if err := a.DB.Find(&categories).Error; err != nil {
		// a.DB.Find(&articles) = SELECT * FROM articles;
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found article"})
		return
	}
	var res []categoriesSuccess
	copier.Copy(&res, &categories)
	ctx.JSON(http.StatusOK, gin.H{"categories": res})
}

func (a *Category) CreatCategory(ctx *gin.Context) {
	// ประกาศตัวแปรที่เก็บ struct ของ createArticle
	var form creatCategory
	// ดัก error ถ้า ctx.ShouldBind(&form) ไม่เท่ากับ null ให้แสดง key เป็น error และแสดงค่าที่ error ออกไป err.Error()
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// ถ้าไม่ error ให้ทำการดึงข้อมูลมาเก็บใน struck ของ article
	var category models.Category
	// copier.Copy(เอาข้อมูล stuck จาก param2 มาใส่ , copy struck ที่ต้องการไปใส่ใน param1)
	copier.Copy(&category, &form)

	// เชื่อมต่อ DB และส่งข้อมูลไปใน DB
	if err := a.DB.Create(&category).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
	categorySuccess := categoriesSuccess{}
	copier.Copy(&categorySuccess, &category)
	ctx.JSON(http.StatusCreated, gin.H{"articles": categorySuccess})
}
