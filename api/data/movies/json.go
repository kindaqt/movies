package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

////////////////////////////////////////
// Persister
//////////////////////////////////////

type JsonStore struct {
	JsonFilePath string
}

func NewJsonStore() *JsonStore {
	return &JsonStore{
		JsonFilePath: "server/data/movies/movies.json",
	}
}

////////////////////////////////////////
// Methods
//////////////////////////////////////

// GetMovies returns all movies
func (p *JsonStore) GetMovies() (interface{}, error) {
	var movies []Movie
	if err := parseJSONFile(&movies, p.JsonFilePath); err != nil {
		return nil, err
	}

	return movies, nil
}

// UpdateWatched updates the watch value value
func (p *JsonStore) UpdateWatched(id string, value bool) error {
	var movies []Movie
	if _, movieIndex, err := findMovie(id, p.JsonFilePath); err != nil {
		movies[movieIndex].Watched = value
	} else {
		return err
	}

	file, err := json.MarshalIndent(movies, "", "")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(p.JsonFilePath, file, 0644); err != nil {
		return err
	}

	return nil
}

///////////////////////////////////////
// Utils
/////////////////////////////////////

func findMovie(id string, jsonFilePath string) (Movie, int, error) {
	var movies []Movie
	if err := parseJSONFile(&movies, jsonFilePath); err != nil {
		return Movie{}, 0, err
	}

	for i, movie := range movies {
		if movie.ID == id {
			return movie, i, nil
		}
	}

	return Movie{}, 0, fmt.Errorf("invalid id")
}

func parseJSONFile(data interface{}, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, &data)
}
