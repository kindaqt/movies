package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	data "github.com/kindaqt/movies/api/data/movies"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

type MockPersister struct {
}

func NewMockStore() data.Store {
	return &MockPersister{}
}

func (p *MockPersister) GetMovies() ([]data.Movie, error) {
	var movies = []data.Movie{{"0", "The Dark Knight", true}}
	return movies, nil
}

func (p *MockPersister) UpdateWatched(id string, value bool) error {
	return nil
}

func (p *MockPersister) DeleteMovie(id string) error {
	return nil
}

func TestGetMoviesHandler(t *testing.T) {
	// Setup HTTP
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Setup Dependencies
	h := Persister{
		Store: NewMockStore(),
	}

	// Run Test
	h.GetMoviesHandler(c)

	// Assertions
	var expectedBody = `[{"id":"0","title":"The Dark Knight","watched":true}]`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, w.Body.String(), expectedBody)
}
