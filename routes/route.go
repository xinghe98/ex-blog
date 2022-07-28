package routes

import (
	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/controller"
	"liandyuan.cn/api/middleware"
)

func RegistRoutes(r *gin.Engine) *gin.Engine {
	root := r.Group("/api/admin")
	root.Use(middleware.Auth(), middleware.RootAuth())
	{
		root.GET("user", controller.Find)               //获取所有用户信息
		root.DELETE("user/:key", controller.DeleteUser) //根据id删除用户
		root.GET("user/profile", controller.Getinfo)    //获取某一个用户的信息
		root.POST("/uploads", controller.Upload)
		root.GET("/uploads/img", controller.Findimg)
		root.DELETE("/uploads/img/:uid", controller.Delimg)
	}
	user := r.Group("/api/")
	{
		user.POST("user", controller.Create)
		user.PUT("user/:key", middleware.Auth(), controller.Updatepassword)
		user.POST("user/login", controller.Login) //用户登录
	}

	return r
}
