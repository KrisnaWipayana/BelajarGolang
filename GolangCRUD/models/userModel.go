package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/KrisnaWipayana/GolangCRUD/config"
	"github.com/KrisnaWipayana/GolangCRUD/entities"
)

type UserModel struct {
	conn *sql.DB
}

func NewUserModel() *UserModel {

	conn, err := config.DBCon() //mengecek koneksi database
	if err != nil {
		panic(err)
	}

	return &UserModel{
		conn: conn, //memasukkan value conn dari function NewUserModel
	}
}

func (p *UserModel) AllUser() ([]entities.User, error) { //membuat function select all dari tabel

	rows, err := p.conn.Query("select * from tb_user") //memilih semua data yang tersimpan di db
	if err != nil {
		return []entities.User{}, err
	}
	defer rows.Close() //menutup akses database

	var dataUser []entities.User
	for rows.Next() {
		var user entities.User
		rows.Scan(&user.Id, &user.Role, &user.Nama, &user.Email, &user.Password) //memindai data untuk row data user

		if user.Role == "1" { //mengubah value jika data di db berupa angka kategori
			user.Role = "Leader"
		} else if user.Role == "2" {
			user.Role = "Co-Leader"
		} else {
			user.Role = "Member"
		}

		dataUser = append(dataUser, user)
	}

	return dataUser, nil
}

func (p *UserModel) Create(user entities.User) bool {

	result, err := p.conn.Exec("insert into tb_user (role, nama, email, password) values(?,?,?,?)", //memasukkan semua data yang dibutuhkan ke dalam tabel user
		user.Role, user.Nama, user.Email, user.Password)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p UserModel) Regist(user entities.User) (int64, error) {

	result, err := p.conn.Exec("insert into tb_user (role, nama, email, password) values(?,?,?,?)", //memasukkan semua data yang dibutuhkan ke dalam tabel user
		user.Role, user.Nama, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId, nil
}

func (p *UserModel) Find(id int64, user *entities.User) error {

	return p.conn.QueryRow("select * from tb_user where id = ?", id).Scan(
		&user.Id,
		&user.Role,
		&user.Nama,
		&user.Email,
		&user.Password)
}

func (p *UserModel) Update(user entities.UserUpdate) error {

	_, err := p.conn.Exec("update tb_user set role = ?, nama = ?, email = ? where id = ?",
		user.Role, user.Nama, user.Email, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *UserModel) Delete(id int64) {

	p.conn.Exec("delete from tb_user where id = ?", id)
}

// func (p *UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

// 	row, err := p.conn.Query("select * from tb_user where "+fieldName+" = ? limit 1", fieldValue)

// 	if err != nil {
// 		return err
// 	}

// 	defer row.Close()

// 	for row.Next() {
// 		row.Scan(&user.Id, &user.Role, &user.Email, &user.Password)
// 	}
// 	return nil
// }

func (p *UserModel) Where(user *entities.User, fieldName, fieldValue string) error {
	// Log untuk memeriksa nilai fieldName dan fieldValue
	fmt.Println("fieldName:", fieldName)
	fmt.Println("fieldValue:", fieldValue)

	query := "select * from tb_user where " + fieldName + " = ? limit 1"
	row, err := p.conn.Query(query, fieldValue)

	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}
	defer row.Close()

	if row.Next() {
		// Log untuk memastikan bahwa data ditemukan
		fmt.Println("User found!")
		err := row.Scan(&user.Id, &user.Role, &user.Nama, &user.Email, &user.Password)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return err
		}
	} else {
		fmt.Println("No user found with the given email")
		return errors.New("user not found")
	}
	return nil
}
