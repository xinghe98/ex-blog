package httpresp

import "github.com/gin-gonic/gin"

func Resp(ctx *gin.Context, status int, code int, msg string) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
	})
}
