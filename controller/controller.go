package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/models"
)

func Pong(ctx *gin.Context) {
	user := models.Test{Name: "笑话", Age: 23}

	models.CreateUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func Find(ctx *gin.Context) {
	all := models.Findall()
	ctx.JSON(http.StatusOK, all)
}
