package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"liandyuan.cn/api/httpresp"
	"liandyuan.cn/api/models"
	"liandyuan.cn/api/util"
)

func Sigup(user *models.User, ctx *gin.Context) {
	if util.ValiUser(user) {
		httpresp.Resp(ctx, http.StatusBadGateway, 502, nil, "用户已经存在")
		return
	}
	if util.Valipassword(user) == false {
		httpresp.Resp(ctx, http.StatusBadGateway, 502, nil, "密码必须大于6位")
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.PassWord), 14)
	user.PassWord = string(bytes)
	times := time.Now().Unix()
	userid := util.MD5(string(times))
	user.UserKey = userid
	err := CreateUser(user)
	if err != nil {
		httpresp.Resp(ctx, http.StatusBadGateway, 502, nil, "请输入正确的信息")
		return
	}
	httpresp.Resp(ctx, http.StatusOK, 200, nil, "注册成功")
}
