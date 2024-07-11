package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
func InitDB(){
	var err error
	DB,err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("error opening database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTables()
}

func CreateTables(){

	CreateUserTables := `Create Table if not exists users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		)`
		
	_,err := DB.Exec(CreateUserTables)

	if err != nil {
		panic("Error creating users table")
	}

	CreateEventTables := `Create Table if not exists events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		location TEXT,
		date_time DATETIME,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
		)`
	_, err = DB.Exec(CreateEventTables)

	if err != nil {
		panic("Error creating events table")
	}
}