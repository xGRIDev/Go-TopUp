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
	createTopupTables := `
	CREATE TABLE IF NOT EXISTS topups (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titlegame TEXT NOT NULL,
		description TEXT NOT NULL,
		price DECIMAL NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	_, err := DB.Exec(createTopupTables)
	if err != nil {
		panic("Could not create Topup Table")
	}
}
