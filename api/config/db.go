package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	return sql.Open("mysql", "disc_user:pass4disc@tcp(172.25.0.3:3306)/disc_db")
}