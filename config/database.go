package config

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	username = "pklvj4y9h37yx9jv"
	password = "bb95vd255c9omqdf"
	hostname = "ffn96u87j5ogvehy.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306"
	dbname   = "hbsvkff78rsechkd"
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
