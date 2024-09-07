package route

import (
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/controller/adminController"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/controller/custController"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/controller/loginController"
	"github.com/KrisnaWipayana/BelajarGolang/GolangFiber/middleware"
	"github.com/gofiber/fiber/v2"
)

// handle routing disini
func AppRoute(c *fiber.App) {

	//middleware group
	auth := c.Group("/admin", middleware.AuthReq)
	order := c.Group("/payment", middleware.AuthReq)

	//Login route
	c.Post("/login", loginController.Login)
	c.Get("/logout", loginController.Logout)
	c.Post("/register", loginController.Register)

	//Customer route
	c.Get("/dashboard", custController.Index)
	order.Get("/details", custController.Order)

	//Admin route
	auth.Get("/user", adminController.ShowUser)
	auth.Get("/user/:id", adminController.GetUser)
	auth.Post("/user", adminController.AddUser)
	auth.Put("/user/:id", adminController.UpdateUser)
	auth.Delete("/user", adminController.DeleteUser)

	//Staff route

}
