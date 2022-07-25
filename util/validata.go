package util

import (
	"liandyuan.cn/api/models"
)

func Valipassword(user *models.User) bool {
	if len(user.PassWord) < 6 {
		return false
	} else {
		return true
	}
}

func ValiUser(user *models.User) bool {
	models.DB.Where("username = ?", user.UserName).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
