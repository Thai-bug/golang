package db

import (
	models "demo/Models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=thai-bug password=12022021 dbname=dev port=9013 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("failed to connect database")
	}

	database.AutoMigrate(&models.User{})
	DB = database
}
