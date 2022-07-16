package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"column:username;not null;unique"`
	PassWord string `json:"password" gorm:"column:password;not null"`
}

//在数据库中建立对应的表结构
func InitModel() {
	DB.AutoMigrate(&User{})
}

//创建用户
func CreateUser(user *User) (err error) {
	err = DB.Create(&user).Error
	return err
}

//查询所有用户
func Findall() (all []*User) {
	DB.Find(&all)
	return all
}

//根据id删除用户
func DeleteUser(id string) {
	DB.Delete(&User{}, id)
}

//根据id更新密码
func Updatepwd(id string, pwd string) {
	DB.Model(&User{}).Where("id=?", id).Update("password", pwd)
}
