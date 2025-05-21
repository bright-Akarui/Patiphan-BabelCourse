```go
############################################ Start.. สร้าง RESTful API ด้วย Gin Framework ############################################
Gin Framwork
	- ทดสอบการใช้งาน
	package main

	import "github.com/gin-gonic/gin"

	func main() {
		r := gin.Default()
		r.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "Hello,World")
		})

		r.Run()
	}
	air github
		- ทำให้เราไม่ต้อง stop server และ start server ขึ้นมาใหม่
		- คำสั่ง go install github.com/cosmtrek/air@latest //go get -u github.com/cosmtrek/air
			- คำสั่ง air
Grouping Routes
	- ตัวอย่างโค้ด
	- การใช้งาน ctx.JSON()
		- ctx.JSON(แจ้งสถานะ http , gin.H{Key: ส่งคืนข้อความ})
		- ตัสอย่างการใช้งาน ctx.JSON(http.StatusOK, gin.H{"articles": articles})
	type article struct {
		Title string
		Body  string
	}

	func Serve(r *gin.Engine) {
		articlesGroup := r.Group("api/v1/articles")
		articles := []article{
			{Title: "text#1", Body: "Body#1"},
			{Title: "text#2", Body: "Body#2"},
			{Title: "text#3", Body: "Body#3"},
			{Title: "text#4", Body: "Body#4"},
			{Title: "text#5", Body: "Body#5"},
		}
		articlesGroup.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"articles": articles})
		})
		articlesGroup.GET("/id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "Hi")
		})
	}
Query และ Params
	Query
		- ยกตัวอย่างเช่น
		//api/v1/articles?limit=3
		articlesGroup.GET("", func(ctx *gin.Context) {
			result := articles
			limit := ctx.Query("limit") // Query
			if limit != "" {
				n, _ := strconv.Atoi(limit)
				result = result[:n]
			}
			ctx.JSON(http.StatusOK, gin.H{"articles": result})
		})
	Params
		- /:id ถ้าเราจะดึงค่าของ id เข้ามาได้ต้องใช้ Params
			- ctx.Params("id")

		articlesGroup.GET("/:id", func(ctx *gin.Context) {
			id, _ := strconv.Atoi(ctx.Param("id"))
			for _, v := range articles {
				if v.ID == uint(id) {
					ctx.JSON(http.StatusOK, gin.H{"articles": v})
					return
				}
			}
			ctx.JSON(http.StatusNotFound, "Notfound")
		})
Form และการ validate ข้อมูลในฟอร์ม
	- เราสามารถสร้างและระบได้ว่าอะไรที่ต้อง required และสามารถกำหนดชื่อที่จะส่งมาจาก json ได้
	type article struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	type createArticle struct {
		Title string `json:"title" binding:"required"`
		Body  string `json:"body" binding:"required"`
	}
	articlesGroup.POST("/post", func(ctx *gin.Context) {
		var form createArticle
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		a := article{
			ID:    uint(len(articles) + 1),
			Title: form.Title,
			Body:  form.Body,
		}

		articles = append(articles, a)
		ctx.JSON(http.StatusCreated, gin.H{"articles": articles})
	})
form data
	- การอัพโหลดไฟล์ การจัดเก็บไฟล์
		ctx.ShouldBind(&form);
		- การสร้าง folder
			- ใช้คำสั่ง os.MkdirAll(ชื่อไฟล์,permission)//permission จะต้องใส่ว่าใครสามารถ read, write, excecute ได้ ดูตาม unix permissions calculator
			- os.MkdirAll("uploads/"+dir,0755)
		- การสร้างตัวแปรเพื่อส่งไฟล์ 
			type createArticle struct {
				Title string `json:"title" binding:"required"`
				Body  string `json:"body" binding:"required"`
				Image *multipart.FileHeader `form:"image" binding:"required"`
			}
		- ขั้นตอนการเก็บไฟล์
				1.Get file
					file,err := ctx.FormFile("image")
				2.Create Path
					path := "uploads/articles/" + strconv.Itoa(int(a.ID))
					os.MkdirAll(path,0775)
				3.Upload file
					filename := path + "/" + file.Filename
					if err := ctx.SaveUploadedFile(file, filename); err != nil {
						return
					}
				4.Attach file to article หรือตัวแปร
					a.Image = "http://127.0.0.1:8080" + filename
		- การที่จะเปิดไฟล์จาก url ได้ ต้องใช้ static ก่อนเสมอ
			//"http://127.0.0.1:8080uploads/articles/6/MainBefore.jpg"
			//r.Static("ภายใต้ path ที่เข้ามา","Folder อะไร")
			r.Static("/upload", "/uploads")
การติดตั้ง env
	กำหนด HOST
		- สร้างไฟล์ .env 
			- HOST=http://127.0.0.1:8080
		- import package "github.com/joho/godotenv"
		- เรียกใช้งาน
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		- วิธีใช้งาน
			a.Image = os.Getenv("HOST") + "/" + filename
	กำหหนด PORT
		- สร้างไฟล์ .env 
			- PORT=5000
		- import package "github.com/joho/godotenv"
		- เรียกใช้งาน
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		- วิธีใช้งาน
			a.Image = os.Getenv("HOST") + "/" + filename
	ควรใส่ gitignore ลงไปเสมอ
		- .env
MVC (Model-view-controller)
	- MVC เป็นรูปแบบการออกแบบซอฟต์แวร์ (Software Design Pattern) ที่แยกส่วนประกอบของโปรแกรมออกเป็น 3 ส่วนหลัก เพื่อให้โค้ด แยกความรับผิดชอบ (Separation of Concerns) และจัดการได้ง่ายขึ้น
	- models
		- ทำหน้าที่
			- จัดการข้อมูล, โครงสร้าง, และติดต่อกับฐานข้อมูล
		- ตัวอย่างในเว็บ
			- ตารางผู้ใช้, ฟังก์ชันดึงข้อมูลจาก DB
	- controller
		- ทำหน้าที่
			- ประสานงานระหว่าง Model และ View	
		- ตัวอย่างในเว็บ
			- รับ request, เรียก Model, ส่งไป View
############################################ END.. สร้าง RESTful API ด้วย Gin Framework ############################################

############################################ START.. การเชื่อมต่อฐานข้อมูลและใช้งาน GORM ############################################
การเชื่อมต่อฐานข้อมูลผ่าน GORM
	- สร้าง folder config
	- สร้างไฟล์ .go ไว้สำหรับเชื่อต่อฐานข้อมูล
	- ดูการเชื่อฐานข้อมูลของแต่และฐานได้ที่ gorm doc
	- ตัวอย่างโค้ด
		package config

		import (
			"log"
			"os"

			"github.com/gin-gonic/gin"
			"gorm.io/driver/postgres"
			"gorm.io/gorm"
			"gorm.io/gorm/logger"
		)

		var db *gorm.DB

		func InitDB() {
			var err error
			var logLevel logger.LogLevel
			if gin.Mode() == gin.DebugMode {
				logLevel = logger.Info
			} else {
				logLevel = logger.Silent
			}

			db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_CONNECTION")), &gorm.Config{
				Logger: logger.Default.LogMode(logLevel),
			})
			if err != nil {
				log.Fatal(err)
			}
		}

		func GetDB() *gorm.DB {
			return db
		}

		func CloseDB() {
			sqlDB, err := db.DB()
			if err != nil {
				log.Println("failed to get sql.DB:", err)
				return
			}
			sqlDB.Close()
		}
		และใน main.go เชื่อมดังนี้
		config.InitDB()
		defer config.CloseDB()
migration
	- gorm.Model จะเป็น stuck ที่เก็บตัวแปรไว้อยู่แล้ว
		type Model struct { // size=88 (0x58)
			ID        uint `gorm:"primarykey"`
			CreatedAt time.Time
			UpdatedAt time.Time
			DeletedAt DeletedAt `gorm:"index"`
		}
	- Title string `gorm:"unique;not null"` //unique คือ ต้องไม่ซ้ำกัน และ ห้ามเป็น null
	- การสร้าง Model
		type Article struct {
			// ID    uint   `json:"id"`
			// Title string `json:"title"`
			// Body  string `json:"body"`
			// Image string `json:"image"`
			gorm.Model //Embed struck ของ gorm model ไว้เพื่อ เอาเข้าตัวแปรมาตั้งต้นไว้
			Title string `gorm:"unique;not null"`
			Excerpt string `gorm:"not null"`
			Body string `gorm:"not null"`
			Image string `gorm:"not null"`
		}
	- การ migration 
		- ยกตัวอย่างถ้าเราอัพโค้ดขึ้นไปครั้งแรกจะเกิดการ migration คือการสร้าง table และ field ขึ้นมาใหม่ ขอแยกเป็น 2 ฝั่ง คือ development และ production
			- development มี table=article(field1:Title) และถูก up ขึ้นไปแล้วทางฝั่ง production 
				- ถ้าฝั่ง develope แก้ไข field table=article(field1:Body) แล้ว up ขึ้นไป ถ้ามีเกิดข้อผิดพลากหลังจาก up แล้วต้องเกิดการ down หรือ rollback กลับด้วย
					- ซึ่งเราจะต้องเขียนดักข้อมูลตรงนี้ไว้ให้เกิดการ rollback
	- การสร้าง migration การสร้างข้อมูลใหม่ในฐานข้อมูล
		- จะใช้ libary ที่ชื่อว่า gormigrate
		- จุดเริ่มต้นสร้าง model เพื่อกำหนด column ใน table ขึ้นมา
			type Article struct {
				// ID    uint   `json:"id"`
				// Title string `json:"title"`
				// Body  string `json:"body"`
				// Image string `json:"image"`
				gorm.Model //Embed struck ของ gorm model ไว้เพื่อ เอาเข้าตัวแปรมาตั้งต้นไว้
				Title string `gorm:"unique;not null"`
				Excerpt string `gorm:"not null"`
				Body string `gorm:"not null"`
				Image string `gorm:"not null"`
			}
		- การสร้าง table และ การ Rollback
			1.สร้าง folder migrations
			2.สร้างไฟล์สำหรับ migration ต่างๆ ยกตัวอย่างเช่น
				m1747737106_create_article_table.go
					package migrations
					import (
						"go-gin/models"
						"github.com/go-gormigrate/gormigrate/v2"
						"gorm.io/gorm"
					)
					func m1747737106CreateArticlesTable() *gormigrate.Migration {
						return &gormigrate.Migration{
							ID: "1747737106", // ID เป็น timestamp
							Migrate: func(tx *gorm.DB) error {
								// ตรวจสอบว่ามี tables ชื่อ articles หรือยังไ ถ้ายังไม่มีจะทำการสร้าง table ขึ้นมา
								return tx.AutoMigrate(&models.Article{}) // จะ return error มาให้เลย ถ้าไม่เท่ากับ null จะ error
							},
							Rollback: func(tx *gorm.DB) error {
								return tx.Migrator().DropTable("articles") // จะ return error มาให้เลย ถ้าไม่เท่ากับ null จะ error
							},
						}
					}
			3.สร้างไฟล์ migration สำหรับ เรียกใช้งาน m1747737106CreateArticlesTable()
				package migrations
				import (
					"go-gin/config"
					"log"
					"github.com/go-gormigrate/gormigrate/v2"
				)
				func Migrate() {
					// สำหรับเรียกใช้การทำงาน
					m := gormigrate.New(
						config.GetDB(),
						gormigrate.DefaultOptions,
						[]*gormigrate.Migration{m1747737106CreateArticlesTable()},
					)
					ดักกรณี error
					if err := m.Migrate(); err != nil {
						log.Fatalf("Migration failed: %v", err)
					}
				}
			4.เรียกใช้ใน main
				migrations.Migrate()
	- CRUD (gorm)
		- แนะนำให้ดูควบคตู๋กับ GORM doc
		- create ตัวอย่างโค้ด และคำอธิบาย
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
		- Read หรือ Query
			- ตัวอย่างโค้ดการทำงาน หาแบบรายไอดี articlesGroup.GET("/:id", articlesController.ArticleFindById)
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
			- ตัสอย่างโค้ดการทำงานแบบ find all
				func (a *Article) FindAll(ctx *gin.Context) {
					var articles []models.Article
					if err := a.DB.Find(&articles).Error; err != nil {
						// a.DB.Find(&articles) = SELECT * FROM users;
						ctx.JSON(http.StatusNotFound, gin.H{"error": err})
						return
					}
					ctx.JSON(http.StatusOK, gin.H{"articles": articles})
				}
		- การสร้างการทำงานของ Pagination
			- การแบ่งข้อมูลการเข้าถึงเป็น page เพื่อลดการประมวลผลในกรณีที่ข้อมูลมีเยอะมากๆ
			- ตัวอย่างการทำงาน
			- 1. การสร้าง
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
					p.query.Limit(limit).Offset(offset).Find(p.records)
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
			2. การใช้งาน
				func (a *Article) FindAll(ctx *gin.Context) {
					var articles []models.Article

					// if err := a.DB.Find(&articles).Error; err != nil {
					// 	// a.DB.Find(&articles) = SELECT * FROM articles;
					// 	ctx.JSON(http.StatusNotFound, gin.H{"error": "not found article"})
					// 	return
					// }
					// res := []articleSuccess{}
					pagination := pagination{ctx: ctx, query: a.DB, records: &articles}
					paging := pagination.pagingResource()
					var res []articleSuccess
					copier.Copy(&res, &articles)
					ctx.JSON(http.StatusOK, gin.H{"articles": articlePaging{Items: res, Paging: paging}})
				}
		- update 
			func (a *Article) UpdateArticle(ctx *gin.Context) {
				// สร้าง struck และกำหนด struck ให้ form
				var form udpateArticle
				if err := ctx.ShouldBind(&form); err != nil {
					ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
					return
				}
				var articles models.Article
				id := ctx.Param("id")
				if err := a.DB.First(&articles, id).Error; err != nil {
					// SELECT * FROM articles WHERE id = ค่าของ id;
					ctx.JSON(http.StatusBadRequest, gin.H{"error": "not found id"})
					return
				}
				if err := a.DB.Model(&articles).Updates(&form).Error; err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"updateError": err.Error()})
					return
				}
				articleSuccess := articleSuccess{}
				copier.Copy(&articleSuccess, &articles)
				ctx.JSON(http.StatusOK, gin.H{"updateAticle": articleSuccess})
			}
		- Delete
			func (a *Article) DeleteArticle(ctx *gin.Context) {
				var article models.Article
				id := ctx.Param("id")

				if err := a.DB.First(&article, id).Error; err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": "not found id"})
					return
				}

				if err := a.DB.Delete(&article).Error; err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete article"})
					return
				}

				ctx.JSON(http.StatusOK, gin.H{"successDelete": "success delete"})
			}
	- Associations
		- จะเหมือนกับการ join table
	- Database Seeding
		- เตรียมข้อมูลสำหัรบการทดสอบบน development
		- สำหรับการสร้าง fake data เข้าไปที่ database มาเพื่อทดสอบการใช้งาน
		- ขัเนตอนการสร้าง
			1. สร้าง package seed
			2. สร้าง seed.go
				func Load(){
					//1.เชื่อมฐานข้อมูล
				}
############################################ END.. การเชื่อมต่อฐานข้อมูลและใช้งาน GORM ############################################