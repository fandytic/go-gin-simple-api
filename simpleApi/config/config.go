package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

func DbConnect() *gorm.DB {
	consStr := "root:@tcp(127.0.0.1:3306)/simple_api?parseTime=true"
	db, err := gorm.Open("mysql", consStr)
	if err != nil {
		log.Fatal("Error when connect db" + consStr + " : " + err.Error())
		return nil
	}
	return db
}