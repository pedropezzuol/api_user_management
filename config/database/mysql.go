package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitializeMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:phpdsdn07@tcp(localhost:3306)/user_management")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
