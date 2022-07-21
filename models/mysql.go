package models

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetMysql() *gorm.DB {
	url := viper.GetString("mysql.url")
	usename := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	fmt.Println(url, usename, password)
	var err error
	database := fmt.Sprintf("%s:%s@tcp%s?charset=utf8mb4&parseTime=True&loc=Local",
		usename, password, url)
	DB, err = gorm.Open(mysql.Open(database), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

//在数据库中建立对应的表结构
func InitModel() {
	DB.AutoMigrate(&User{})
}
