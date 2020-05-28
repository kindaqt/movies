package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	jsonFilePath = "server/data/movies/movies.json"
)

type Movie struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Watched bool   `json:"watched"`
}

// GetMovies returns all movies
func GetMovies() (interface{}, error) {

	var movies []Movie
	if err := parseJSONFile(&movies, jsonFilePath); err != nil {
		return nil, err
	}

	return movies, nil
}

// UpdateWatched updates the watch value value
func UpdateWatched(id string, value bool) error {
	var movies []Movie
	if _, movieIndex, err := findMovie(id); err != nil {
		movies[movieIndex].Watched = value
	} else {
		return err
	}

	file, err := json.MarshalIndent(movies, "", "")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(jsonFilePath, file, 0644); err != nil {
		return err
	}

	return nil
}

///////////////////////////////////////
// Utils
/////////////////////////////////////

// returns index, error
func findMovie(id string) (Movie, int, error) {
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
