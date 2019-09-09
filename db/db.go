package db

import (
	"database/sql"
	"fmt"

	//Postgress import
	_ "github.com/lib/pq"
)

// Item struct
type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//data array
var data []Item

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

//AddItem function
func AddItem(item *Item) error {
	sqlState := `
	INSERT INTO item (title, description)
	VALUES ($1, $2);`

	_, err := db.Exec(sqlState, &item.Title, &item.Description)

	if err != nil {
		return err
	}

	return nil
}

//FindItem func
func FindItem(item Item) (*[]Item, error) {
	var data []Item

	var rows *sql.Rows
	var err error

	sqlState := `
		SELECT * 
		FROM item 
		WHERE
		title LIKE $1 and
		description LIKE $2`

	if item.ID != 0 {
		sqlState += ` and id = $3`
		rows, err = db.Query(sqlState, item.Title+"%", item.Description+"%", item.ID)

	} else {
		rows, err = db.Query(sqlState, item.Title+"%", item.Description+"%")

	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := Item{}
		err = rows.Scan(
			&item.Title,
			&item.Description,
			&item.ID,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)

	}

	return &data, nil

}

//GetAllItems func
func GetAllItems() *[]Item {
	var data []Item
	rows, err := db.Query(
		`
			SELECT 
			id,
			title,
			description FROM item;
		`)
	if err != nil {
		fmt.Println("Errr", err)
		return &data
	}

	defer rows.Close()

	for rows.Next() {
		item := Item{}

		err = rows.Scan(
			&item.ID,
			&item.Title,
			&item.Description,
		)

		if err != nil {
			return &data
		}
		data = append(data, item)
	}
	return &data
}

//GetLenght func
func GetLenght() int {
	return len(data)
}
