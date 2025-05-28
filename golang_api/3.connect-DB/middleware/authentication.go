package middleware

import (
	"go-gin/config"
	"go-gin/models"
	"log"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

var identityKey = "sub"

func Authenticate() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		// secret key ต้องทราบว่าเอามาจากไหนและมีค่าเป็นอะไร
		// Key: []byte(os.Getenv("SECRET_KEY")),
		Key: []byte(os.Getenv("SECRET_KEY")),

		IdentityKey: identityKey,

		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		//ค่าของ user
		IdentityHandler: func(ctx *gin.Context) interface{} {
			var user models.User
			claims := jwt.ExtractClaims(ctx)
			id := claims[identityKey]

			db := config.GetDB()
			if err := db.First(&user, uint(id.(float64))).Error; err != nil {
				return nil
			}
			return &user
		},

		// login นำข้อมูลของ user ที่ต้องการ login มาตรวจสอบ
		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			var form login
			var user models.User

			if err := ctx.ShouldBindJSON(&form); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			db := config.GetDB()
			if err := db.Where("email = ?", form.Email).First(&user).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},

		// แปลงจาก user เป็น payload(sub) และแปลงเป็น json wen token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				claims := jwt.MapClaims{
					identityKey: v.ID,
				}

				return claims
			}
			return jwt.MapClaims{}
		},
		// ถ้า login ไม่ผ่าน ต้องตรวจสอบข้อผิดพลาด
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
