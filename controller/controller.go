package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	all := models.Findall()
	ctx.JSON(http.StatusOK, all)
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
	models.DeleteUser(id)
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
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	models.Updatepwd(id, string(bytes))
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// 客户查询用户信息
func Getinfo(ctx *gin.Context) {
	user, _ := ctx.Get("userinfo")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": user},
	})
}
