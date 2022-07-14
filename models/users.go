package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string  `json:"username"`
	Age      int64   `json:"age"`
	PassWord string  `json:"password"`
	Doc      *string `json:"doc"`
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
