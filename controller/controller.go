package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/models"
)

func Create(ctx *gin.Context) {
	//user := models.Test{Name: "笑话", Age: 23}
	var user models.User
	ctx.BindJSON(&user)
	models.CreateUser(&user)
	fmt.Println(user)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func Find(ctx *gin.Context) {
	all := models.Findall()
	ctx.JSON(http.StatusOK, all)
}
