package main

import (
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"liandyuan.cn/api/models"
	"liandyuan.cn/api/routes"
)

type Test struct {
	gorm.Model
	Name string
	Age  int64
	Doc  *string
}

func main() {
	r := gin.Default()
	db := models.ConnetMysql()
	//db.AutoMigrate(&Test{})
	str := "this is a test"
	test := Test{Name: "xinghe", Age: 18, Doc: &str}
	db.Create(&test)
	r = routes.UserRoute(r)
	r.Run()
}
