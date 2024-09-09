package middleware

import (
	"fmt"

	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/session"
	"github.com/gofiber/fiber/v2"
)

func AuthReq(allowedRoles ...int) fiber.Handler {

	return func(c *fiber.Ctx) error {
		//mendapatkan session dari request
		sess, err := session.Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"auth": "gagal mendapatkan session",
			})
		}

		token := sess.Get("jwt")
		fmt.Println("token didapat : ", token)
		fmt.Println("----------------------------------------------------------------------------")

		if token == nil {
			fmt.Println("session tidak ditemukan")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"auth": "session tidak ada, silahkan login",
			})
		}

		role := sess.Get("role")
		fmt.Println("role didapat : ", role)

		if role == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"role-session": "gagal mendapatkan role",
			})
		}

		roleInt, ok := role.(int)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"role-format": "invalid role format",
			})
		}
		for _, allowedRole := range allowedRoles {
			if roleInt == allowedRole {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"middleware": "akses role dilarang",
		})
	}
}
