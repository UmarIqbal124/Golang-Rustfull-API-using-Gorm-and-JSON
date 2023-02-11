package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	urlDSN   = "root:@tcp(127.0.0.1:3306)/apicheck?parseTime=true"
	err      error
)

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection not established successfully! ", err)
	}
	Database.AutoMigrate(&Employee{})

}
