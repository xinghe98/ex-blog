package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "gorm.io/driver/mysql"
	"liandyuan.cn/api/middleware"
	"liandyuan.cn/api/models"
	"liandyuan.cn/api/router"
)

func InitConfig() {
	worddir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(worddir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

		fmt.Println(err)
		panic(err)
	}
}
func main() {
	InitConfig()
	r := gin.Default()
	r.Use(middleware.Cors())
	models.ConnetMysql()
	models.InitModel()
	r = router.RegistRoutes(r)
	r.Run(":3000")
}
