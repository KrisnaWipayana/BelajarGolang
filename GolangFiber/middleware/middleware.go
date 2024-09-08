package middleware

import (
	"fmt"

	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/session"
	"github.com/gofiber/fiber/v2"
)

func AuthReq(c *fiber.Ctx) error {

	//mendapatkan session dari request
	sess, err := session.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mendapatkan session",
		})
	}

	token := sess.Get("jwt")
	fmt.Println("token: ", token)

	if token == nil {
		fmt.Println("session tidak ditemukan") //nyangkut disini
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "session tidak ada, silahkan login",
		})
	}
	return c.Next()
}
