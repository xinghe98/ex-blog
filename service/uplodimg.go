package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"liandyuan.cn/api/models"
)

func Uploadimg(c *gin.Context) bool {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		uuid := uuid.New()
		key := uuid.String()
		dst := fmt.Sprintf("../vue-ft/public/uploads/%s", key)
		// 上传文件到指定的目录
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return false
		}
		var index int = 1
		if key, _ := c.Get("key"); key == "root" {
			index = 0
		}
		imginfo := models.Imageinfo{Uid: key, Imgname: file.Filename, Imgurl: fmt.Sprintf("/uploads/%s", key), Index: index}
		models.DB.Create(&imginfo)
	}
	return true
}

//查询所有图片
func Findimages() (all []*models.Imageinfo) {
	models.DB.Find(&all)
	return all
}

//根据id删除图片
func Deleteimg(uid string) {
	var img models.Imageinfo
	models.DB.Unscoped().Where("uid=?", uid).Delete(&img)
}
