package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"liandyuan.cn/api/httpresp"
	"liandyuan.cn/api/service"
)

func Upload(c *gin.Context) {
	// Multipart form
	if service.Uploadimg(c) == false {

		httpresp.Resp(c, http.StatusBadRequest, 502, nil, "上传失败")
		return
	}
	httpresp.Resp(c, http.StatusOK, 200, nil, "上传成功")
}
