package config

import (
	"database/sql"
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

func ConnectDB() {
	var err error
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
	DB, err = sql.Open("mysql", database)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to Database")
}
