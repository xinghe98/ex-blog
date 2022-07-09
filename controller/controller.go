package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
