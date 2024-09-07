package main

import (
	"fmt"
	"net/http"

	authController "github.com/KrisnaWipayana/GolangCRUD/controller/authController"
	"github.com/KrisnaWipayana/GolangCRUD/controller/userController"
)

func main() {

	http.HandleFunc("/", authController.Login)
	http.HandleFunc("/login", authController.Login)
	http.HandleFunc("/logout", authController.Logout)
	http.HandleFunc("/register", authController.Register)
	http.HandleFunc("/index", authController.Index)
	http.HandleFunc("/user", userController.User)
	http.HandleFunc("/user/index", userController.User)
	http.HandleFunc("/user/add", userController.Add)
	http.HandleFunc("/user/edit", userController.Edit)
	http.HandleFunc("/user/delete", userController.Delete)

	http.ListenAndServe(":3000", nil)
	fmt.Println("Server berjalan di port 3000, buka localhost:3000 untuk mengakses web")
}
