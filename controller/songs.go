package controller

import (
	"example/ProyectoLTV/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSongs(c *gin.Context) {
	queryResult, errFound := dbInstance.Query(" SELECT genres.name AS genre, songs.length, songs.song, songs.artist " +
		" FROM songs " +
		" INNER JOIN genres " +
		" ON genres.id = songs.ID ")

	if errFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem retrieving songs data"})
		return
	}
	songs := make([]model.Song, 0)

	for queryResult.Next() {
		song := model.Song{}
		errFound = queryResult.Scan(&song.Genre, &song.Length, &song.Song, &song.Artist)
		if errFound != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Problem retrieving songs data"})
		}
		songs = append(songs, song)

	}

	errFound = queryResult.Err()

	if errFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem retrieving songs data"})
	}

	if songs == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem retrieving songs data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}

func GetSongsByArtist(c *gin.Context) {
	queryResult, errFound := dbInstance.Query("SELECT genres.name AS genre, songs.length, songs.song, songs.artist " +
		" FROM songs " +
		" INNER JOIN genres " +
		" ON genres.id = songs.ID " + " WHERE songs.artist = '" + c.Param("artist") + "'")

	if errFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem retrieving songs data"})
		return
	}
	songs := make([]model.Song, 0)

	for queryResult.Next() {
		song := model.Song{}
		errFound = queryResult.Scan(&song.Genre, &song.Length, &song.Song, &song.Artist)
		if errFound != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Problem retrieving songs data"})
			return
		}
		songs = append(songs, song)

	}

	errFound = queryResult.Err()
	if errFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem retrieving songs data"})
		return
	}

	if songs == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problem retrieving songs data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": songs})
}
