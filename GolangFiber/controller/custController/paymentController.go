package custController

import "github.com/gofiber/fiber/v2"

func Order(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "order coming soon xixixi",
	})
}
