package controller

import (
	"example/ProyectoLTV/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// QueryString used in some functions to select elements with join
var errorMessage = "Problem Retrieving Songs"
var selectSongsQuery = " SELECT genres.name AS genre, songs.length, songs.song, songs.artist " +
	" FROM songs " +
	" INNER JOIN genres " +
	" ON genres.id = songs.ID "

// This will send an HTTP indicating an error ocurred
func handleErrors(queryError error, c *gin.Context) {
	if queryError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
	}
	return
}

func GetAllSongs(c *gin.Context) {
	queryResult, queryError := dbInstance.Query(selectSongsQuery)

	handleErrors(queryError, c)
	songs := make([]model.Song, 0)

	for queryResult.Next() {
		song := model.Song{}
		queryError = queryResult.Scan(&song.Genre, &song.Length, &song.Song, &song.Artist)
		handleErrors(queryError, c)
		songs = append(songs, song)

	}

	queryError = queryResult.Err()

	handleErrors(queryError, c)

	if songs == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}
func GetSongsByArtist(c *gin.Context) {
	queryResult, queryError := dbInstance.Query(selectSongsQuery + " WHERE songs.artist = '" + c.Param("artist") + "'")

	handleErrors(queryError, c)
	songs := make([]model.Song, 0)

	for queryResult.Next() {
		song := model.Song{}
		queryError = queryResult.Scan(&song.Genre, &song.Length, &song.Song, &song.Artist)
		handleErrors(queryError, c)
		songs = append(songs, song)

	}

	queryError = queryResult.Err()
	handleErrors(queryError, c)

	if songs == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}
func GetSongsBySongName(c *gin.Context) {
	queryResult, queryError := dbInstance.Query(selectSongsQuery + " WHERE songs.song = '" + c.Param("nameSearched") + "'")

	handleErrors(queryError, c)
	songs := make([]model.Song, 0)

	for queryResult.Next() {
		song := model.Song{}
		queryError = queryResult.Scan(&song.Genre, &song.Length, &song.Song, &song.Artist)
		handleErrors(queryError, c)
		songs = append(songs, song)

	}

	queryError = queryResult.Err()
	handleErrors(queryError, c)

	if songs == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}
func GetSongsByGenre(c *gin.Context) {
	queryResult, queryError := dbInstance.Query(selectSongsQuery + " WHERE genres.name = '" + c.Param("genre") + "'")

	handleErrors(queryError, c)
	songs := make([]model.Song, 0)

	for queryResult.Next() {
		song := model.Song{}
		queryError = queryResult.Scan(&song.Genre, &song.Length, &song.Song, &song.Artist)
		handleErrors(queryError, c)
		songs = append(songs, song)
	}

	queryError = queryResult.Err()
	handleErrors(queryError, c)

	if songs == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}
func GetSongsByLength(c *gin.Context) {
	var minLength = c.Query("minLen")
	var maxLength = c.Query("maxLen")

	//Check if the params exist
	if maxLength == "" && minLength == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	queryResult, queryError := dbInstance.Query(selectSongsQuery + " WHERE songs.length > " + minLength + " AND songs.length < " +
		maxLength +
		" ORDER BY songs.length")

	handleErrors(queryError, c)
	songs := make([]model.Song, 0)

	for queryResult.Next() {
		song := model.Song{}
		queryError = queryResult.Scan(&song.Genre, &song.Length, &song.Song, &song.Artist)
		handleErrors(queryError, c)
		songs = append(songs, song)
	}

	queryError = queryResult.Err()
	handleErrors(queryError, c)

	if songs == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}
