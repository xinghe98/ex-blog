package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"column:username;not null;unique"`
	PassWord string `json:"password" gorm:"column:password;not null"`
}

func InitModel() {
	DB.AutoMigrate(&User{})
}

func CreateUser(user *User) {

	DB.Create(&user)
}

func Findall() (all []*User) {
	DB.Find(&all)
	return all
}
