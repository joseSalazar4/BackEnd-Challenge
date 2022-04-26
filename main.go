package main

import (
	"example/ProyectoLTV/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func SetupRouter() *gin.Engine {
	// Generate the router instance
	router := gin.Default()

	// Connect to database and generate an instance to use
	controller.ConnectToDatabase()

	// These are the routes created containing main tasks and extra credit
	router.GET("/songs", controller.GetAllSongs)
	router.GET("/songsByLength", controller.GetSongsByLength)
	router.GET("/songsByGenre/:genre", controller.GetSongsByGenre)
	router.GET("/songsByArtist/:artist", controller.GetSongsByArtist)
	router.GET("/songsSortedByGenre", controller.GetSongsStatsSortedByGenre)
	router.GET("/songsBySongName/:nameSearched", controller.GetSongsBySongName)

	return router
}

func main() {
	// Call the func to connect to database and set routes!
	router := SetupRouter()

	// Start Listening! Port: 8080
	router.Run()

}
