package models

import (
	"fmt"

	"events.com/db"
	"events.com/utils"
)

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	// Save the user to the database
	u.Password = utils.HashPassword(u.Password)
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

func GetUser(email string) (User,error) {

	row := db.DB.QueryRow("SELECT * FROM users WHERE email = ?", email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}