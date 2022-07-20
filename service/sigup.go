package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"liandyuan.cn/api/httpresp"
	"liandyuan.cn/api/models"
	"liandyuan.cn/api/util"
)

func Sigup(user *models.User, ctx *gin.Context) {
	if util.ValiUser(user) {
		httpresp.Resp(ctx, http.StatusBadGateway, 502, "用户已经存在")
		return
	}
	if util.Valipassword(user) == false {
		httpresp.Resp(ctx, http.StatusBadGateway, 502, "密码必须大于11位")
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.PassWord), 14)
	user.PassWord = string(bytes)
	err := models.CreateUser(user)
	if err != nil {
		httpresp.Resp(ctx, http.StatusBadGateway, 502, "请输入正确的信息")
		return
	}
	httpresp.Resp(ctx, http.StatusOK, 200, "注册成功")
}
