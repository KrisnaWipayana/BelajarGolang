package custController

import (
	"log"

	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/database"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/model/entities"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	//start show layanan
	var layanan []entities.Layanan
	err := database.DB.Find(&layanan).Error
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"error": "gagal mengambil data layanan",
		})
	}
	//end show layanan

	//start show kamar
	var kamar []entities.Kamar
	err = database.DB.Find(&kamar).Error
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"error": "gagal mengambil data kamar",
		})
	}

	//end show kamar
	response := fiber.Map{
		"layanan": layanan,
		"kamar":   kamar,
	}
	return c.JSON(response)
}
