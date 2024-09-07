package adminController

import (
	"encoding/json"
	"log"

	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/database"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/model/entities"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/req"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ShowUser(c *fiber.Ctx) error {

	var user []entities.User
	err := database.DB.Find(&user).Error
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"error": "gagal mengambil data layanan",
		})
	}
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {

	var user []entities.User
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validator": "ID tidak boleh kosong",
		})
	}
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat menemukan user",
		})
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data di dapat",
		"data":    user,
	})
}

func AddUser(c *fiber.Ctx) error {

	user := new(req.CreateUser)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validator": "Gagal input",
			"error":     err.Error(),
		})
	}

	hashPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal hash password",
		})
	}

	newUser := entities.User{
		Role:     user.Role,
		Nama:     user.Nama,
		Email:    user.Email,
		Password: string(hashPw),
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memasukkan data user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil create user",
		"data":    newUser,
	})
}

func UpdateUser(c *fiber.Ctx) error {

	var user entities.User
	id := c.Params("id")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if database.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "gagal mengupdate user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil mengupdate user",
	})
}

func DeleteUser(c *fiber.Ctx) error {

	var input struct {
		Id json.Number
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	id, err := input.Id.Int64()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID tidak valid"})
	}

	if database.DB.Delete(&entities.User{}, id).RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "tidak dapat menghapus user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil menghapus user",
	})
}
