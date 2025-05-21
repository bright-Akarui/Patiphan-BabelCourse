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

type pagination struct {
	ctx     *gin.Context
	query   *gorm.DB
	records interface{}
}

func (p *pagination) pagingResource() *pagingResult {
	// 1. Get limit, page ? limit=10&page=2
	// ctx.Query("limit")
	// กรณที่ไม่ได้ระบุค่า limit เข้ามา จะ default เป็น 12
	limit, _ := strconv.Atoi(p.ctx.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(p.ctx.DefaultQuery("page", "1"))
	ch := make(chan int)
	// 2. Count records มีข้อมูลทั้งหมดกี่ตัว
	go p.countRecord(ch)
	// 3. Find Records
	// สูตรหาค่าของ offset
	offset := (page - 1) * limit
	p.query.Preload("Category").Limit(limit).Offset(offset).Find(p.records)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// 4. total nextPage จำนวนของ page ทั้งหมด
	count := <-ch
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
		Count:     int(count),
		TotalPage: totalPage,
	}
}
func (p *pagination) countRecord(ch chan int) {
	// 2. Count records มีข้อมูลทั้งหมดกี่ตัว
	var count int64
	// SELECT COUNT(*) FROM records;
	p.query.Model(p.records).Count(&count)
	ch <- int(count)
}
