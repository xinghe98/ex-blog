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

//根据id删除用户
func DeleteUser(id string) {
	models.DB.Delete(&models.User{}, id)
}

//根据id更新密码
func Updatepwd(id string, pwd string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	models.DB.Model(&models.User{}).Where("id=?", id).Update("password", string(bytes))
}
