package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	responseWriter := httptest.NewRecorder()
	r.ServeHTTP(responseWriter, request)
	return responseWriter
}

const expectedResponseGetSongs string = "{\"data\":[{\"length\":189,\"artist\":\"424\",\"song\":\"Gala\",\"genre\":\"Country\"},{\"length\":246,\"artist\":\"Colornoise\",\"song\":\"Amalie\",\"genre\":\"Rap\"},{\"length\":165,\"artist\":\"Los Waldners\",\"song\":\"Horacio\",\"genre\":\"Classical\"},{\"length\":235,\"artist\":\"Chubby Checker\",\"song\":\"The Twist\",\"genre\":\"Indie Rock\"},{\"length\":167,\"artist\":\"Santana\",\"song\":\"Smooth\",\"genre\":\"Noise Rock\"},{\"length\":245,\"artist\":\"Bobby Darin\",\"song\":\"Mack the Knife\",\"genre\":\"Latin Pop Rock\"},{\"length\":237,\"artist\":\"LeAnn Rhimes\",\"song\":\"How Do I Live\",\"genre\":\"Classic Rock\"},{\"length\":189,\"artist\":\"LMFAO\",\"song\":\"Party Rock Anthem\",\"genre\":\"Pop\"}]}"
const expectedResponseGetByLen string = "{\"data\":[{\"length\":159,\"song\":\"Macarena\"},{\"length\":162,\"song\":\"Hey Jude\"},{\"length\":165,\"song\":\"Horacio\"},{\"length\":167,\"song\":\"Smooth\"}]}"
const expectedResponseGetByGenre string = "{\"data\":[{\"length\":189,\"artist\":\"LMFAO\",\"song\":\"Party Rock Anthem\",\"genre\":\"Pop\"}]}"
const expectedResponseGetByArtist string = "{\"data\":[{\"length\":237,\"artist\":\"LeAnn Rhimes\",\"song\":\"How Do I Live\",\"genre\":\"Classic Rock\"}]}"
const expectedResponseGetBySongName string = "{\"data\":[{\"length\":167,\"artist\":\"Santana\",\"song\":\"Smooth\",\"genre\":\"Noise Rock\"}]}"
const expectedResponseGetSortedByGenre string = "{\"data\":[{\"length\":237,\"song\":\"1\",\"genre\":\"Classic Rock\"},{\"length\":165,\"song\":\"1\",\"genre\":\"Classical\"},{\"length\":189,\"song\":\"1\",\"genre\":\"Country\"},{\"length\":235,\"song\":\"1\",\"genre\":\"Indie Rock\"},{\"length\":245,\"song\":\"1\",\"genre\":\"Latin Pop Rock\"},{\"length\":167,\"song\":\"1\",\"genre\":\"Noise Rock\"},{\"length\":189,\"song\":\"1\",\"genre\":\"Pop\"},{\"length\":246,\"song\":\"1\",\"genre\":\"Rap\"}]}"

func TestGetSongs(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"data": expectedResponseGetSongs,
	}

	router := SetupRouter()
	// Perform a GET request with the handler.
	w := performRequest(router, "GET", "/songs")

	// Chek if responseCode is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Compare the JSON response to what the API should reply with
	value := w.Body.String()
	assert.Equal(t, body["data"], value)
}

func TestGetSongsByLength(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"data": expectedResponseGetByLen,
	}

	router := SetupRouter()
	// Perform a GET request with the handler.
	w := performRequest(router, "GET", "/songsByLength?minLen=1&maxLen=180")

	// Chek if responseCode is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Compare the JSON response to what the API should reply with
	value := w.Body.String()
	assert.Equal(t, body["data"], value)
}

func TestGetSongsByGenre(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"data": expectedResponseGetByGenre,
	}

	router := SetupRouter()
	// Perform a GET request with the handler.
	w := performRequest(router, "GET", "/songsByGenre/Pop")

	// Chek if responseCode is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Compare the JSON response to what the API should reply with
	value := w.Body.String()
	assert.Equal(t, body["data"], value)
}

func TestGetSongsByArtist(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"data": expectedResponseGetByArtist,
	}

	router := SetupRouter()
	// Perform a GET request with the handler.
	w := performRequest(router, "GET", "/songsByArtist/LeAnn Rhimes")

	// Chek if responseCode is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Compare the JSON response to what the API should reply with
	value := w.Body.String()
	assert.Equal(t, body["data"], value)
}

func TestGetSongsSortedByGenre(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"data": expectedResponseGetSortedByGenre,
	}

	router := SetupRouter()
	// Perform a GET request with the handler.
	w := performRequest(router, "GET", "/songsSortedByGenre")

	// Chek if responseCode is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Compare the JSON response to what the API should reply with
	value := w.Body.String()
	assert.Equal(t, body["data"], value)
}

func TestGetSongsBySongName(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"data": expectedResponseGetBySongName,
	}

	router := SetupRouter()
	// Perform a GET request with the handler.
	w := performRequest(router, "GET", "/songsBySongName/Smooth")

	// Chek if responseCode is 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Compare the JSON response to what the API should reply with
	value := w.Body.String()
	assert.Equal(t, body["data"], value)
}
