package dao

import (
	"book_spider/config"
	"book_spider/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	fmt.Println("Start Dao Init")

	dbURL := config.GetConfig("mysql.url", "127.0.0.1:3306")
	dbUserName := config.GetConfig("mysql.userName", "root")
	dbPassword := config.GetConfig("mysql.password", "123456")
	dbName := config.GetConfig("mysql.databaseName", "ship")

	fmt.Println("Ulr:", dbURL)
	dsn := dbUserName + ":" + dbPassword + "@tcp(" + dbURL + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Link Mysql Eerror", err.Error())
	}
	initTable()
}

func initTable() {
	db.AutoMigrate(&model.Flight{}, &model.Ship{})
}
