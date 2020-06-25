package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMoviesSuccess(t *testing.T) {
	store, err := NewJsonStore("./movies.json")
	assert.NoError(t, err, "NewJsonStore() should not return an error.")

	movies, err := store.GetMovies()
	assert.NoError(t, err, "GetMovies should not return an error")

	expected := []Movie{
		{"0", "Shawshank Redemption", false},
		{"1", "The Godfather", false},
		{"2", "The Godfather: Part II", false},
		{"3", "The Dark Knight", false},
	}
	assert.Equal(t, expected, movies)
}
