package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "root@tcp(127.0.0.1.3306)/smartcanteen?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("cannot connect to database")
	}

	fmt.Println("connected to database")

	DB = conn
	DB.AutoMigrate(models.seller{}, models.buyers{}, models.product)
}
