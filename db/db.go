package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() {
	database = initDatabase()
}

func GetDB() *gorm.DB {
	return database
}

func initDatabase() *gorm.DB {
	host := "127.0.0.1"
	port := "3306"
	user := "root"
	pwd := "12345678"
	name := "MSHD"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
