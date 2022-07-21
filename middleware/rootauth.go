package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if key, _ := ctx.Get("key"); key != "root" {
			ctx.JSON(http.StatusForbidden, gin.H{"code": 403, "msg": "无权限"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
