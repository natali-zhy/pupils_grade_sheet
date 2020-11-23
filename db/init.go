package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func New(user string, password string, database string) {
	dataSourceName := "user=" + user + " password=" + password + " dbname=" + database + " sslmode=disable"

	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Errorf("db: %w", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Errorf("db: %w", err)
	}

	DB = db
}
