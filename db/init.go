package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
)

var DB *sql.DB

func InitDB() {
	conn, err := sql.Open("postgres", "user=postgres password=1111 dbname=students sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB OK")
	}

	DB = conn
}

