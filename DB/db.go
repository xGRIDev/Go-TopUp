package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func DBinit() {
	var err error
	DB, err = sql.Open("sqlite3", "topups.db")

	if err != nil {
		panic("Could Connect To Database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
}

func createTable() {

	// USERS-CREATE-TABLE
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Cannot create Users Table.")
	}

	createTopupTables := `
	CREATE TABLE IF NOT EXISTS topups (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titlegame TEXT NOT NULL,
		description TEXT NOT NULL,
		price DECIMAL NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createTopupTables)
	if err != nil {
		panic("Could not create Topup Table")
	}
}
