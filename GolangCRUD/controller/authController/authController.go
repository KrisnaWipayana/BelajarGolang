package authController

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/KrisnaWipayana/GolangCRUD/config"
	"github.com/KrisnaWipayana/GolangCRUD/entities"
	"github.com/KrisnaWipayana/GolangCRUD/models"
	"golang.org/x/crypto/bcrypt"
)

// func Index(w http.ResponseWriter, r *http.Request) {

// 	temp, _ := template.ParseFiles("views/dashboard/login.html")
// 	temp.Execute(w, nil)
// }

type UserInput struct {
	Email    string
	Password string
}

var userModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"nama": session.Values["nama"],
			}

			temp, _ := template.ParseFiles("views/dashboard/index.html")
			temp.Execute(w, data)
		}
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	fmt.Println(session)

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/dashboard/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {

		r.ParseForm()
		userInput := &UserInput{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		var user entities.User
		userModel.Where(&user, "email", userInput.Email)

		var message error

		if user.Email == "" {
			message = errors.New("email salah")
		} else {

			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
			if errPassword != nil {
				message = errors.New("password salah") //masalah utama error
			}
		}

		if message != nil {

			data := map[string]interface{}{
				"error": message,
			}

			temp, _ := template.ParseFiles("views/dashboard/login.html")
			temp.Execute(w, data)
		} else {

			session.Values["loggedIn"] = true
			session.Values["role"] = user.Role
			session.Values["nama"] = user.Nama
			session.Values["email"] = user.Email

			session.Save(r, w)

			http.Redirect(w, r, "/index", http.StatusSeeOther)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	//menghapus session
	session.Options.MaxAge = -1
	session.Save(r, w)

	fmt.Println(session)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		//menampilkan views
		temp, _ := template.ParseFiles("views/dashboard/register.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		//melakuka proses registrasi
		r.ParseForm()

		user := entities.User{
			Role:      r.Form.Get("role"),
			Nama:      r.Form.Get("nama"),
			Email:     r.Form.Get("email"),
			Password:  r.Form.Get("password"),
			Cpassword: r.Form.Get("cpassword"),
		}

		errorMessage := make(map[string]interface{})

		if user.Role == "" {
			errorMessage["Role"] = "Role harus diisi"
		}
		if user.Nama == "" {
			errorMessage["Nama"] = "Nama harus diisi"
		}
		if user.Email == "" {
			errorMessage["Email"] = "Email harus diisi"
		}
		if user.Password == "" {
			errorMessage["Password"] = "Password harus diisi"
		}
		if user.Cpassword == "" {
			errorMessage["Cpassword"] = "Konfirmasi password harus diisi"
		} else {
			if user.Cpassword != user.Password {
				errorMessage["Cpassword"] = "Konfirmasi password tidak cocok"
			}
		}
		if len(errorMessage) > 0 { //validasi form gagal

			data := map[string]interface{}{
				"validate": errorMessage,
			}
			temp, _ := template.ParseFiles("views/dashboard/register.html")
			temp.Execute(w, data)
		} else {

			//hash pw
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			//insert ke db
			_, err := userModel.Regist(user)

			var message string

			if err != nil {
				message = "Proses registrasi gagal: " + message
			} else {
				message = "Registrasi berhasil"
			}
			data := map[string]interface{}{
				"alert": message,
			}
			temp, _ := template.ParseFiles("views/dashboard/register.html")
			temp.Execute(w, data)
		}
	}
}
