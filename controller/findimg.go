package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/httpresp"
	"liandyuan.cn/api/service"
)

func Findimg(ctx *gin.Context) {
	all := service.Findimages()
	httpresp.Resp(ctx, http.StatusOK, 200, all, "查询成功")
}

func Delimg(ctx *gin.Context) {
	key, ok := ctx.Params.Get("uid")
	if !ok {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 502,
			"msg":  "请求无效",
		})
		return
	}
	service.Deleteimg(key)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
