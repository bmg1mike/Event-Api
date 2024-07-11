package models

import (
	"fmt"

	"events.com/db"
)

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	// Save the user to the database
	rowsAffected, err := db.DB.Exec("INSERT INTO users(email, password) VALUES(?,?)", u.Email, u.Password)
	if err != nil {
		return err
	}

	value, err := rowsAffected.RowsAffected()
	
	if err != nil{
		return err
	}
	
	fmt.Println("the number of rows affected are" + string(value))
	return nil
}