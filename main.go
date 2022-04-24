package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connection to Database is initializaed
	db, err := sql.Open("sqlite3", "songsDatabaseLTV.db")
	if err != nil {
		fmt.Println("there is no conecction")
	}
	res, err := db.Query("SELECT song FROM songs")
	fmt.Println(err)
	var song string
	for res.Next() {
		res.Scan(&song)
		fmt.Println(song)
	}
}
