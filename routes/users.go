package routes

import (
	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/controller"
)

func UserRoute(r *gin.Engine) *gin.Engine {
	user := r.Group("/api/")
	{
		user.POST("user", controller.Create)
		user.GET("user", controller.Find)
		user.DELETE("user/:id", controller.DeleteUser)
		user.PUT("user/:id", controller.Updatepassword)
	}
	return r
}
