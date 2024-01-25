package config

import (
	"database/sql"
	"log"

	"github.com/pedropezzuol/api_test/config/database"
)

var db *sql.DB

func Init() error {
	var err error

	db, err = database.InitializeMysql()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetMysql() *sql.DB {
	return db
}
