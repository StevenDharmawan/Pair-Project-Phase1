package config

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	username = ""
	password = ""
	hostname = ""
	dbname   = ""
)

func ConnectDB() error {
	var err error
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	DB, err = sql.Open("mysql", database)
	if err != nil {
		return errors.New("failed connect to database")
	}
	fmt.Println("Connected to Database")
	return nil
}
