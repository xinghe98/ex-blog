// 用于user的增删改查
package service

import (
	"golang.org/x/crypto/bcrypt"
	"liandyuan.cn/api/models"
)

//创建用户
func CreateUser(user *models.User) (err error) {
	err = models.DB.Create(&user).Error
	return err
}

//查询所有用户
func Findall() (all []*models.User) {
	models.DB.Find(&all)
	return all
}

//根据userkey删除用户
func DeleteUser(key string) {
	var user models.User
	models.DB.Unscoped().Where("userkey=?", key).Delete(&user)
}

//根据id更新密码
func Updatepwd(key string, pwd string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	models.DB.Model(&models.User{}).Where("userkey=?", key).Update("password", string(bytes))
}
