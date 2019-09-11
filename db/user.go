package db

import (
	"database/sql"
	"fmt"
)

//User struct
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Items    []Item `json:"items"`
}

//AddUser func
func AddUser(user *User) error {

	sqlState := `
	INSERT INTO app_user (username, full_name)
	VALUES ($1, $2);`

	_, err := db.Exec(sqlState, user.Username, user.FullName)
	if err != nil {
		return err
	}
	return nil
}

//FindUser func
func FindUser(user *User) (*[]User, error) {
	var response []User
	var rows *sql.Rows
	var itemRows *sql.Rows
	var err error

	sqlItem := `
		SELECT * 
		FROM item
		WHERE
		user_id = $1;
	`
	sqlState := `
	SELECT * 
	FROM app_user 
	WHERE
	username LIKE $1 and
	full_name LIKE $2`

	if user.ID != 0 {
		sqlState += ` and id = $3`
		rows, err = db.Query(sqlState, user.Username+"%", user.FullName+"%", user.ID)
	} else {
		rows, err = db.Query(sqlState, user.Username+"%", user.FullName+"%")
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.FullName,
		)
		if err != nil {
			return nil, err
		}
		itemRows, err = db.Query(sqlItem, user.ID)
		var items []Item
		for itemRows.Next() {
			item := Item{}
			err = itemRows.Scan(
				&item.Title,
				&item.Description,
				&item.ID,
				&item.UserID,
			)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
		user.Items = items
		response = append(response, user)
	}
	return &response, nil
}

//DeleteUser func
func DeleteUser(id int) error {
	sqlState := `DELETE FROM item
	WHERE user_id = $1;`
	_, err := db.Exec(sqlState, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	sqlState = `
	DELETE FROM app_user
	WHERE id = $1;`
	_, err = db.Exec(sqlState, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//UpdateUser func
func UpdateUser(user *User) error {
	sqlState := `
	UPDATE app_user
	SET username = $1,
	full_name = $2
	WHERE id = $3;`
	_, err := db.Exec(sqlState, user.Username, user.FullName, user.ID)
	if err != nil {
		return err
	}
	return nil
}
