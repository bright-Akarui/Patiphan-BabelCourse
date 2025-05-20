package controllers

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type pagingResult struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	PrevPage  int `json:"prevPage"`
	NextPage  int `json:"nextPage"`
	Count     int `json:"count"`
	TotalPage int `json:"totalPage"`
}

func pagingResource(ctx *gin.Context, query *gorm.DB, records interface{}) *pagingResult {
	// 1. Get limit, page ? limit=10&page=2
	// ctx.Query("limit")
	// กรณที่ไม่ได้ระบุค่า limit เข้ามา จะ default เป็น 12
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "12"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "12"))

	// 2. Count records มีข้อมูลทั้งหมดกี่ตัว
	var count int64
	// SELECT COUNT(*) FROM records;
	query.Model(records).Count(&count)

	// 3. Find Records
	// สูตรหาค่าของ offset
	offset := (page - 1) * limit
	query.Limit(limit).Offset(offset).Find(&records)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// 4. total nextPage จำนวนของ page ทั้งหมด
	totalPage := int(math.Ceil(float64(count) / float64(limit)))

	// 5. Find nextPage
	var nextPage int
	if page == totalPage {
		nextPage = totalPage
	} else {
		nextPage = page + 1
	}

	// 6. create pagingResult
	return &pagingResult{
		Page:      page,
		Limit:     limit,
		PrevPage:  page - 1,
		NextPage:  nextPage,
		TotalPage: totalPage,
	}
}
