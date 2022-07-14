package models

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetMysql() *gorm.DB {
	var err error
	dsn := "root:xinghe@tcp(109.236.63.125:3306)/usersinfo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}
