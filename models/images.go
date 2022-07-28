package models

import "gorm.io/gorm"

type Imageinfo struct {
	gorm.Model
	Uid     string `json:"uid" gorm:"column:uid"`
	Imgname string `json:"imgname" gorm:"column:imgname"`
	Imgurl  string `json:"imgurl" gorm:"column:imgurl"`
	Index   int    `json:"index" gorm:"column:index"`
}
