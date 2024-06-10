package config

import (
	// "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// dsn := "root:1234@tcp(127.0.0.1:3306)/Book_Store?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:1234@tcp(127.0.0.1:3306)/Book_Store?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
