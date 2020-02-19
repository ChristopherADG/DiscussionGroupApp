package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	DB, err := sql.Open("mysql", "disc_user:pass4disc@tcp(172.25.0.3:3306)/disc_db")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You are connected to database aurant_db...")
}