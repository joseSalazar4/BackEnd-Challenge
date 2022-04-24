package main

import (
	"example/ProyectoLTV/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	// Connection to Database is initializaed
	controller.ConnectToDatabase()
	router.GET("/songs", controller.GetAllSongs)
	router.Run()

}
