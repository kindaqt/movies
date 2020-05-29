package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	data "github.com/kindaqt/movies/api/data/movies"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type MockPersister struct {
	JsonFilePath string
}

func (p *MockPersister) GetMovies() (interface{}, error) {
	var movies = []data.Movie{data.Movie{"0", "The Dark Knight", true}}
	return movies, nil
}

func (p *MockPersister) UpdateWatched(id string, value bool) error {
	return nil
}

func TestGetMoviesHandler(t *testing.T) {
	// Setup HTTP
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Setup Dependencies
	client := &data.Client{
		Store: &MockPersister{
			JsonFilePath: "./test/data/movies.json",
		},
	}
	dataPersister = *client

	// Run Test
	GetMoviesHandler(c)

	// Assertions
	var expectedBody = `[{"id":"0","title":"The Dark Knight","watched":true}]`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), expectedBody)
}
