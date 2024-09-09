package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	const conn = "root:@tcp(127.0.0.1:3306)/db_api_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	DSN := conn

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Database gagal terkoneksi")
	}
	fmt.Println("Terkoneksi ke database...")
}
