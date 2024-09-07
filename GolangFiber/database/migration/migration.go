package migration

import (
	"fmt"

	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/database"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/model/entities"
)

func Migrate() {

	database.DB.AutoMigrate(&entities.User{})
	database.DB.AutoMigrate(&entities.Kamar{})
	database.DB.AutoMigrate(&entities.Foto{})
	database.DB.AutoMigrate(&entities.Layanan{})
	database.DB.AutoMigrate(&entities.No_Kamar{})
	database.DB.AutoMigrate(&entities.Pemesanan{})

	fmt.Println("Migrating...")
}
