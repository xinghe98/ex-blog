package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserKey  string `json:"userkey" gorm:"column:userkey;not null;unique;"`
	UserName string `json:"username" gorm:"column:username;not null;unique"`
	PassWord string `json:"password" gorm:"column:password;not null"`
}
