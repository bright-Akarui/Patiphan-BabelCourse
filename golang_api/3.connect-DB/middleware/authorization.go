package middleware

import (
	"go-gin/models"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.Get("sub")
		if !ok {
			// การใช้ AbortWithStatusJSON ซึ่งจะทำให้ handler ตัวต่อไปไม่ทำงาน
			// ตัวอย่าง categoriesGroup.POST("/post", authorization, categoriesController.CreatCategory) มันจะทำ authorization จนจบและไม่ทำตัวถัดไป
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// user, _ := ctx.Get("sub")
		// if user.(*models.User).Role != "Admin" {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 	return
		// }
		enforcer := casbin.NewEnforcer("config/acl_model.conf", "config/policy.csv")
		u, ok := user.(*models.User)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
			return
		}
		ok = enforcer.Enforce(u, ctx.Request.URL.Path, ctx.Request.Method)

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you are not allowed to access"})
			return
		}

		ctx.Next()
	}
}
