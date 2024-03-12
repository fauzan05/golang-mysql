package golangmysql

import (
	"database/sql"
	"time"
)


func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10) // set default koneksi saat idle
	db.SetMaxOpenConns(100) // set maksimum koneksi 
	db.SetConnMaxIdleTime(5 * time.Minute) // jika lebih dari 5 menit tidak ada yang menggunakan, maka akan kembali ke default (10 koneksi)
	db.SetConnMaxLifetime(60 * time.Minute)
	
	return db
}