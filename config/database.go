package config

// db yang kita pake disini adalah sql. jadi install dulu package driver golang sql
// https://github.com/go-sql-driver/mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// penulisan Database menggunakan kapital agar golang tau bahwa variable ini bisa diimport ke package lain
var Database *sql.DB

func ConnectDB(username string, password string, dbname string) {

	// sql.Open() bertugas membuka database yang diinginkan.
	// sql.Open() return 2 nilai, yaitu koneksi dan error.
	// tipe data koneksi yang direturn bertipe *sql.DB. silahkan cek sendiri tipenya.
	db, err := sql.Open("mysql", username+":"+password+"@/"+dbname+"?parseTime=true")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected!")
	}

	// intinya, variable Database ini menyimpan nilai koneksi database.
	Database = db
}
