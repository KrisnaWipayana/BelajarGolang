package staffController

import (
	"log"

	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/database"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/model/entities"
	"github.com/gofiber/fiber/v2"
)

func ShowKamar(c *fiber.Ctx) error {

	var kamar []entities.Kamar
	err := database.DB.Find(&kamar).Error
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"staff-controller": "gagal mengambil data kamar",
		})
	}

	response := (fiber.Map{
		"kamar": kamar,
	})
	return c.JSON(response)
}

func ShowLayanan(c *fiber.Ctx) error {

	var layanan []entities.Layanan
	err := database.DB.Find(&layanan).Error
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"staff-controller" : "gagal mengambil data layanan"
		})
	}
	response := fiber.Map{
		"layanan": layanan,
	}
	return c.JSON(response)
}
