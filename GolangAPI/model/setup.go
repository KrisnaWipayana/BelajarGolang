package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_go_api"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Foto{})
	database.AutoMigrate(&Kamar{})
	database.AutoMigrate(&No_Kamar{})
	database.AutoMigrate(&Layanan{})
	database.AutoMigrate(&Pemesanan{})

	DB = database
}
