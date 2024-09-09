package route

import (
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/controller/adminController"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/controller/custController"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/controller/loginController"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/controller/staffController"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/middleware"
	"github.com/gofiber/fiber/v2"
)

// handle routing disini
func AppRoute(c *fiber.App) {

	//middleware group
	admin := c.Group("/admin", middleware.AuthReq)
	staff := c.Group("/staff", middleware.AuthReq)
	order := c.Group("/payment", middleware.AuthReq)

	//Login route
	c.Post("/login", loginController.Login)
	c.Get("/logout", loginController.Logout)
	c.Post("/register", loginController.Register)

	//Customer route
	c.Get("/dashboard", custController.Index)
	order.Get("/details", custController.Order)

	//Admin route
	admin.Get("/user", adminController.ShowUser)
	admin.Get("/user/:id", adminController.GetUser)
	admin.Post("/user", adminController.AddUser)
	admin.Put("/user/:id", adminController.UpdateUser)
	admin.Delete("/user", adminController.DeleteUser)

	//Staff route
	staff.Get("/layanan", staffController.ShowLayanan)
	staff.Get("/kamar", staffController.ShowKamar)
}
