package conf

import (
	"database/sql"
	"fmt"
	"log"
)

// DB Database instance
var DB *sql.DB

// Connect database function
func Connect() error {
	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_HOST"), Config("DB_PORT"), Config("DB_NAME")))
	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Connection Opened to Database")
	return nil
}
