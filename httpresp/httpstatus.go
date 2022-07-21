package httpresp

import "github.com/gin-gonic/gin"

func Resp(ctx *gin.Context, status int, code int, data interface{}, msg string) {
	ctx.JSON(status, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
