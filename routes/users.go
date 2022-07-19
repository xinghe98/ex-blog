package routes

import (
	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/controller"
	"liandyuan.cn/api/middleware"
)

func UserRoute(r *gin.Engine) *gin.Engine {
	user := r.Group("/api/")
	{
		user.POST("user", controller.Create)
		user.GET("user", controller.Find)
		user.DELETE("user/:id", controller.DeleteUser)
		user.PUT("user/:id", controller.Updatepassword)
		user.POST("user/login", controller.Login)
		user.GET("user/profile", middleware.Auth(), controller.Getinfo)
	}
	return r
}
