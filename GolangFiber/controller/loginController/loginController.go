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
			"input-error": "invalid input",
		})
	}

	// Find email
	var user entities.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"input-error": "gagal find email",
		})
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"login-controller-Login": "gagal verif password",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claim := token.Claims.(jwt.MapClaims)
	claim["id"] = user.ID
	claim["email"] = user.Email
	claim["role"] = user.Role
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("admin123"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"login-controller-Login": "gagal mendapat token rahasia :(",
		})
	}

	sess, err := session.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"login-controller-Login": "gagal mengakses session",
		})
	}

	// Set session data
	sess.Set("jwt", t)
	sess.Set("role", user.Role)
	// fmt.Println("JWT sebelum disimpan:", t) -- sudah di test (DONE)

	if err := sess.Save(); err != nil {
		fmt.Println("Gagal menyimpan sesi:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"login-controller-Login": "Gagal menyimpan session",
		})
	}

	// unsolved : tidak bisa memanggil token jwt
	// fmt.Println("Session disimpan dengan token:", sess.Get("jwt"))

	return c.JSON(fiber.Map{
		"message":       "berhasil login",
		"token rahasia": t,
		"role":          claim["role"],
	})
}

func Logout(c *fiber.Ctx) error {

	c.ClearCookie("jwt") //clear cookie yang sudah tersambung sebelumnya
	c.ClearCookie("role")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"login-controller-Logout": "berhasil logout, byebye",
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
			"login-controller-Register": "Gagal registrasi",
			"error":                     err.Error(),
		})
	}

	//hashing pw biar keren
	hashPw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"login-controller-Register": "gagal hash password",
		})
	}

	newUser := entities.User{
		Nama:     user.Nama,
		Email:    user.Email,
		Password: string(hashPw),
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"login-controller-Register": "Gagal registrasi user",
		})
	}
	return c.JSON(fiber.Map{
		"login-controller-Register": "berhasil registrasi",
		"data":                      newUser,
	})
}
