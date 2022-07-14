package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Name string
	Age  int64
	Doc  *string
}

func CreateUser(test *Test) {
	DB.Create(&test)
}

func Findall() (all []*Test) {
	DB.Find(&all)
	return all
}
