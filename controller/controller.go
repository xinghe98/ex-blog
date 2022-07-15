package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/models"
	"liandyuan.cn/api/service"
)

// 新建用户
func Create(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	if service.ValiUser(user) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 502,
			"msg":  "该用户名已存在",
		})
		return
	}
	if service.Valipassword(&user) == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 502,
			"msg":  "密码必须大于6位",
		})
		return
	}
	models.CreateUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

// 查询所有用户
func Find(ctx *gin.Context) {
	all := models.Findall()
	ctx.JSON(http.StatusOK, all)
}
