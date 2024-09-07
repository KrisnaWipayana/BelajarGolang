package loginController

import (
	"fmt"
	"time"

	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/database"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/model/entities"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/req"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/session"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {

	var input req.LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid input",
		})
	}

	//find email
	var user entities.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"email": "gagal find email",
		})
	}

	//verif pw
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"password": "gagal verif password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claim := token.Claims.(jwt.MapClaims)
	claim["id"] = user.ID                                //mengambil id user
	claim["email"] = user.Email                          //mengambil email user
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix() //mendapatkan berdurasi 24 jam

	//membuat secrey key paling rahasia sejagat raya
	t, err := token.SignedString([]byte("admin123"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapat token rahasia :(",
		})
	}

	sess, err := session.Store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mengakses session",
		})
	}

	sess.Set("jwt", t)
	fmt.Println("Session sebelum disimpan:", sess.Get("jwt"))

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menyimpan session",
		})
	}

	return c.JSON(fiber.Map{
		"message":       "berhasil login anjay",
		"token rahasia": t, // jika berhasil login maka mengembalikan pesan json beserta token
	})
}

func Logout(c *fiber.Ctx) error {

	c.ClearCookie("jwt") //clear cookie yang sudah tersambung sebelumnya

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil logout, byebye",
	})
}

func Register(c *fiber.Ctx) error {

	user := new(req.NewUser)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal registrasi",
			"error":   err.Error(),
		})
	}

	//hashing pw biar keren
	hashPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal hash password",
		})
	}

	newUser := entities.User{
		Nama:     user.Nama,
		Email:    user.Email,
		Password: string(hashPw),
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal registrasi user",
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil registrasi",
		"data":    newUser,
	})
}
