package main

import (
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
	"liandyuan.cn/api/models"
	"liandyuan.cn/api/routes"
)

func main() {
	r := gin.Default()
	models.ConnetMysql()
	models.InitModel()
	r = routes.UserRoute(r)
	r.Run(":3000")
}
