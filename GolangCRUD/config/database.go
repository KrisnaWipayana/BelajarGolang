package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBCon() (*sql.DB, error) {
	DbDriver := "mysql"
	DbUser := "root"
	DbPass := ""
	DbName := "db_belajargo"

	db, err := sql.Open(DbDriver, DbUser+":"+DbPass+"@/"+DbName)
	return db, err
}
