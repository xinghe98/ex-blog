package main

import (
	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/routes"
)

func main() {
	r := gin.Default()
	r = routes.Route(r)
	r.Run()
}
