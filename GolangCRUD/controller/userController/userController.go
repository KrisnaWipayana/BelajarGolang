package userController

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/KrisnaWipayana/GolangCRUD/config"
	"github.com/KrisnaWipayana/GolangCRUD/entities"
	"github.com/KrisnaWipayana/GolangCRUD/libraries"
	"github.com/KrisnaWipayana/GolangCRUD/models"
)

var validation = libraries.NewValidation() //memanggil function validasi di dalam library validasi
var userModel = models.NewUserModel()

func User(response http.ResponseWriter, request *http.Request) {

	session, _ := config.Store.Get(request, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(response, request, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(response, request, "/login", http.StatusSeeOther)
		} else {
			user, _ := userModel.AllUser()

			data := map[string]interface{}{
				"user": user,
			}
			temp, err := template.ParseFiles("views/user/user.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(response, data)
		}
	}
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/user/add.html") //menampilkan views
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var user entities.User
		user.Role = request.Form.Get("role")
		user.Nama = request.Form.Get("nama")
		user.Email = request.Form.Get("email")
		user.Password = request.Form.Get("password")

		var data = make(map[string]interface{}) //membuat map dari variable data

		vErrors := validation.Struct(user)

		if vErrors != nil {
			data["user"] = user //mengembalikan variable user yg diatas ke views
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data berhasil disimpan"
			userModel.Create(user)
		}

		temp, _ := template.ParseFiles("views/user/add.html")
		temp.Execute(response, data) //variabel data di map pesan dimasukkan ke dalam parameter
		//temp untuk memasukkan pesan ke dalam interface (temp = template)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		idStr := queryString.Get("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(response, "Invalid user ID", http.StatusBadRequest)
			return
		} //masalah id user not found :)

		// Mendapatkan data user berdasarkan ID
		var user entities.User
		err = userModel.Find(id, &user)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				http.Error(response, "User not found", http.StatusNotFound)
				return
			}
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Mempersiapkan data untuk dikirim ke template
		data := map[string]interface{}{
			"user": user,
		}

		// Memparsing dan mengeksekusi template untuk menampilkan data
		temp, err := template.ParseFiles("views/user/edit.html")
		if err != nil {
			panic(err)
		}
		fmt.Println("User id: ", id)
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var user entities.UserUpdate

		user.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		user.Role = request.Form.Get("role")
		user.Nama = request.Form.Get("nama")
		user.Email = request.Form.Get("email")
		// user.Password = request.Form.Get("password")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(user)

		if vErrors != nil {
			data["pesan"] = "Gagal edit"
			fmt.Println(vErrors)
			data["user"] = user
			data["validation"] = vErrors
		} else {
			fmt.Println("Berhasil di edit")
			data["pesan"] = "Data berhasil diedit"
			userModel.Update(user)
		}

		temp, _ := template.ParseFiles("views/user/edit.html")
		temp.Execute(response, data)
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	userModel.Delete(id)

	http.Redirect(response, request, "/user", http.StatusSeeOther)
}
