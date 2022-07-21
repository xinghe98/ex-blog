package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"liandyuan.cn/api/models"
	"liandyuan.cn/api/util"
)

func Login(ctx *gin.Context) {
	//获取数据
	var loguser models.User
	ctx.BindJSON(&loguser)
	//验证数据
	var user models.User
	models.DB.Where("username=?", loguser.UserName).First(&user)
	if user.UserKey == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(loguser.PassWord))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	//发放token

	token, err := util.GenToken(user.UserName, user.UserKey, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 500,
			"msg":  "token生成错误",
		})
		return
	}
	//返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}
