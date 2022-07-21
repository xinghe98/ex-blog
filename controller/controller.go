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
	key, ok := ctx.Params.Get("key")
	if !ok {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 502,
			"msg":  "请求无效",
		})
		return
	}
	service.DeleteUser(key)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

// 更新用户信息
func Updatepassword(ctx *gin.Context) {
	postkey, ok := ctx.Params.Get("key")
	if !ok {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": 502,
			"msg":  "请求无效",
		})
		return
	}
	key, _ := ctx.Get("key")
	if key != postkey {
		httpresp.Resp(ctx, http.StatusForbidden, 403, nil, "请求错误")
		return
	}
	pwd := ctx.PostForm("password")
	service.Updatepwd(postkey, pwd)
	httpresp.Resp(ctx, http.StatusOK, 200, nil, "修改成功")
}

// 客户查询用户信息
func Getinfo(ctx *gin.Context) {
	var user models.User
	userkey, _ := ctx.Get("key")
	models.DB.Where("userkey =?", userkey).Find(&user)
	httpresp.Resp(ctx, http.StatusOK, 200, gin.H{"user": user}, "")
}
