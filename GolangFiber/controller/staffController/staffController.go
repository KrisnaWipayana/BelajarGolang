package staffController

import (
	"encoding/json"
	"log"

	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/database"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/model/entities"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/req"
	"github.com/go-playground/validator/v10"
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

func GetKamar(c *fiber.Ctx) error {

	var kamar []entities.Kamar
	idKamar := c.Params("id")
	if idKamar == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validator": "ID kamar tidak boleh kosong",
		})
	}
	if err := database.DB.Where("id = ?", idKamar).First(&kamar).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat menemukan kamar",
		})
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data di dapat",
		"data":    kamar,
	})
}

func AddKamar(c *fiber.Ctx) error {

	kamar := new(req.CreateKamar)
	if err := c.BodyParser(kamar); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(kamar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validator": "Gagal input",
			"error":     err.Error(),
		})
	}

	newKamar := entities.Kamar{
		NamaKamar:      kamar.NamaKamar,
		Harga:          kamar.Harga,
		DeskripsiKamar: kamar.DeskripsiKamar,
	}
	if err := database.DB.Create(&newKamar).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memasukkan data kamar",
		})
	}
	return c.JSON(fiber.Map{
		"message": " berhasil create kamar",
		"data":    newKamar,
	})
}

func UpdateKamar(c *fiber.Ctx) error {

	var kamar entities.Kamar
	idKamar := c.Params("id")

	if err := c.BodyParser(&kamar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if database.DB.Model(&kamar).Where("id = ?", idKamar).Updates(&kamar).RowsAffected == 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal megupdate kamar",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengupdate user",
	})
}

func DeleteKamar(c *fiber.Ctx) error {

	var input struct {
		Id json.Number
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id, err := input.Id.Int64()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID kamar tidak valid",
		})
	}
	if database.DB.Delete(&entities.Kamar{}, id).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat menghapus kamar",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil menghapus kamar",
	})
}

func ShowLayanan(c *fiber.Ctx) error {

	var layanan []entities.Layanan
	err := database.DB.Find(&layanan).Error
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"staff-controller": "gagal mengambil data layanan",
		})
	}
	response := fiber.Map{
		"layanan": layanan,
	}
	return c.JSON(response)
}

func GetLayanan(c *fiber.Ctx) error {

	var layanan []entities.Layanan
	idLayanan := c.Params("id")
	if idLayanan == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validator": "ID layanan tidak boleh kosong",
		})
	}
	if err := database.DB.Where("id = ?", idLayanan).First(&layanan).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat menemukan layanan",
		})
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data di dapat",
		"data":    layanan,
	})
}

func AddLayanan(c *fiber.Ctx) error {

	layanan := new(req.NewLayanan)
	if err := c.BodyParser(layanan); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(layanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validator": "Gagal input",
			"error":     err.Error(),
		})
	}

	newLayanan := entities.Layanan{
		NamaLayanan: layanan.NamaLayanan,
		Deskripsi:   layanan.Deskripsi,
	}
	if err := database.DB.Create(&newLayanan).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memasukkan data layanan",
		})
	}
	return c.JSON(fiber.Map{
		"message": " berhasil create layanan",
		"data":    newLayanan,
	})
}

func UpdateLayanan(c *fiber.Ctx) error {

	var layanan entities.Layanan
	idLayanan := c.Params("id")

	if err := c.BodyParser(&layanan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if database.DB.Model(&layanan).Where("id = ?", idLayanan).Updates(&layanan).RowsAffected == 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal megupdate layanan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengupdate layanan",
	})
}

func DeleteLayanan(c *fiber.Ctx) error {

	var input struct {
		Id json.Number
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id, err := input.Id.Int64()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID layanan tidak valid",
		})
	}
	if database.DB.Delete(&entities.Layanan{}, id).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat menghapus layanan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil menghapus layanan",
	})
}
