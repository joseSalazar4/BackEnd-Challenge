package main

import (
	"example/ProyectoLTV/controller"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connection to Database is initializaed
	controller.ConnectToDatabase()

}
