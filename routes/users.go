package routes

import (
	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/controller"
)

func UserRoute(r *gin.Engine) *gin.Engine {
	user := r.Group("/api/user")
	{
		user.GET("ping", controller.Pong)
	}
	return r
}
