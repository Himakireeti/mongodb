package database

import (
	"fmt"
	"models/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErr(err error, error_message string) {
	if err != nil {
		fmt.Println(err, " ", error_message)
	}
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	Database_setup()
	return DB
}

func Database_setup() {

	dsn := "root:@tcp(127.0.0.1:3306)/partyholic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErr(err, "Database Created")

	db.Debug().AutoMigrate(&models.Addresses{}, &models.Users{}, &models.Parties{}, &models.CancelledParties{})
	DB = db

}
