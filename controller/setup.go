package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Instance to be used by others
var dbInstance *sql.DB

func ConnectToDatabase() {
	// Connection to Database is initializaed
	db, err := sql.Open("sqlite3", "songsDatabaseLTV.db")

	// Check if there's an error connecting to db. Halt exec
	if err != nil {
		fmt.Println("Error when connecting to database")
		panic(err)
	}
	fmt.Println("Connection to DB estabislhed!")
	dbInstance = db
}
