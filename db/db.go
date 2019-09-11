package db

import (
	"database/sql"
	"fmt"

	//Postgress import
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "tajna"
	dbname   = "db_1"
)

var db *sql.DB

//ConnectDB func
func ConnectDB() {
	psqlInfo := fmt.Sprintf("port=%d user=%s "+
		"password=%s dbname=%s host=%s sslmode=disable",
		port, user, password, dbname, host)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DBErr:::", err)
		fmt.Println("Reconnecting to DB...")
		ConnectDB()
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("PINGErr:::", err)
		fmt.Println("Pinging...")
		ConnectDB()
	}
	fmt.Println("Database connected")
}
