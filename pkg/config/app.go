package config

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:AllDogsLive1!/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
