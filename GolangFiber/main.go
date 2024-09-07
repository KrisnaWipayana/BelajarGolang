package main

import (
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/database"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/database/migration"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/route"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/session"
	"github.com/gofiber/fiber/v2"
)

func main() {

	session.Session()    // set session jadi func global
	database.ConnectDB() //koneksi database
	migration.Migrate()  //auto migrate
	app := fiber.New()   //buat instance fiber

	app.Use(func(c *fiber.Ctx) error {
		sess, err := session.Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message-main": "error mengakses session",
			})
		}
		defer sess.Save()
		return c.Next()
	})

	// app.Use(middleware.AuthReq)

	route.AppRoute(app) //direct ke package route, function AppRoute dgn parameter app
	app.Listen(":3000") //server dijalanin di port 3000
}
