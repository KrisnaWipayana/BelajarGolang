package main

import (
	"github.com/KrisnaWipayana/BelajarGO/GolangAPI/controller/adminController"
	"github.com/KrisnaWipayana/BelajarGO/GolangAPI/controller/authController"
	"github.com/KrisnaWipayana/BelajarGO/GolangAPI/controller/indexController"
	"github.com/KrisnaWipayana/BelajarGO/GolangAPI/model"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	model.ConnectDB()

	//auth route
	r.GET("/", authController.Login)
	r.POST("/", authController.Login)
	r.GET("/logout", authController.Logout)

	//customer route
	r.GET("/index", indexController.Index)
	r.GET("/kamar/:id", indexController.Kamar)
	r.GET("/pesanan", indexController.Pemesanan)
	r.POST("/pesanan", indexController.Pemesanan)

	//admin route
	r.GET("api/user", adminController.User)
	r.GET("api/user/:id", adminController.Show)
	r.POST("api/user", adminController.Add)
	r.PUT("api/user/:id", adminController.Update)
	r.DELETE("api/user", adminController.Delete)

	r.Run()
}
