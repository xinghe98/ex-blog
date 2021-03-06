package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/util"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//从请求头获取token
		tokenstr := ctx.GetHeader("Authorization")
		if tokenstr == "" || !strings.HasPrefix(tokenstr, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请先登录",
			})
			ctx.Abort()
			return
		}
		tokenstr = tokenstr[7:]
		token, claims, err := util.ParseToken(tokenstr)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusForbidden, gin.H{
				"code": 403,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		userkey := claims.UserKey
		if userkey == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		ctx.Set("key", userkey)
	}
}
