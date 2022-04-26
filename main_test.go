package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

const expectedResponse string = "{\"data\":[{\"length\":189,\"artist\":\"424\",\"song\":\"Gala\",\"genre\":\"Country\"},{\"length\":246,\"artist\":\"Colornoise\",\"song\":\"Amalie\",\"genre\":\"Rap\"},{\"length\":165,\"artist\":\"Los Waldners\",\"song\":\"Horacio\",\"genre\":\"Classical\"},{\"length\":235,\"artist\":\"Chubby Checker\",\"song\":\"The Twist\",\"genre\":\"Indie Rock\"},{\"length\":167,\"artist\":\"Santana\",\"song\":\"Smooth\",\"genre\":\"Noise Rock\"},{\"length\":245,\"artist\":\"Bobby Darin\",\"song\":\"Mack the Knife\",\"genre\":\"Latin Pop Rock\"},{\"length\":237,\"artist\":\"LeAnn Rhimes\",\"song\":\"How Do I Live\",\"genre\":\"Classic Rock\"},{\"length\":189,\"artist\":\"LMFAO\",\"song\":\"Party Rock Anthem\",\"genre\":\"Pop\"}]}"

func TestHelloWorld(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"data": expectedResponse,
	}
	// Grab our router
	router := SetupRouter()
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/songs")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	value := w.Body.String()
	assert.Equal(t, body["data"], value)
}
