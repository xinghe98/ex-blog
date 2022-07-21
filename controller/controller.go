package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/httpresp"
	"liandyuan.cn/api/models"
	"liandyuan.cn/api/service"
)

// 新建用户
func Create(ctx *gin.Context) {
	var user *models.User
	ctx.BindJSON(&user)
	service.Sigup(user, ctx)
}

// 查询所有用户
func Find(ctx *gin.Context) {
	all := service.Findall()
	httpresp.Resp(ctx, http.StatusOK, 200, all, "查询成功")
}

//删除用户
func DeleteUser(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 502,
			"msg":  "请求无效",
		})
		return
	}
	service.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

// 更新用户信息
func Updatepassword(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 502,
			"msg":  "请求无效",
		})
		return
	}
	pwd := ctx.PostForm("password")
	service.Updatepwd(id, pwd)
	httpresp.Resp(ctx, http.StatusOK, 200, nil, "修改成功")
}

// 客户查询用户信息
func Getinfo(ctx *gin.Context) {
	user, _ := ctx.Get("userinfo")
	httpresp.Resp(ctx, http.StatusOK, 200, gin.H{"user": user}, "")
}
