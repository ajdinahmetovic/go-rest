package db

import (
	"database/sql"
	"fmt"

	//Postgress import
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "tajna"
	dbname   = "db_1"
)

var db *sql.DB

//ConnectDB func
func ConnectDB() {

	psqlInfo := fmt.Sprintf("port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DB says: ", err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
