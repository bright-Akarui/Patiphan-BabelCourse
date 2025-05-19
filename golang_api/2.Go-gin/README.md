```go
############################################ Start.. สร้าง RESTful API ด้วย Gin Framewprk ############################################
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