package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

////////////////////////////////////////
// Persister
//////////////////////////////////////

type JsonStore struct {
	JsonFilePath string
}

func NewJsonStore(filePath string) Store {
	log.Println("Json: NewJsonStore")
	absPath, _ := filepath.Abs(filePath)
	return &JsonStore{
		JsonFilePath: absPath,
	}
}

////////////////////////////////////////
// Methods
//////////////////////////////////////

// GetMovies returns all movies
func (p *JsonStore) GetMovies() ([]Movie, error) {
	fmt.Println("Json: GetMovies")

	var movies []Movie
	if err := parseJSONFile(&movies, p.JsonFilePath); err != nil {
		return nil, err
	}

	return movies, nil
}

// UpdateWatched updates the watch value value
func (p *JsonStore) UpdateWatched(id string, value bool) error {
	fmt.Println("Json: UpdateWatched")
	movies, err := p.GetMovies()
	if err != nil {
		return err
	}

	_, movieIndex, err := findMovie(id, movies)
	if err != nil {
		return err
	} else {
		movies[movieIndex].Watched = value
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

func findMovie(id string, movies []Movie) (Movie, int, error) {
	fmt.Println("Finding movie...")
	for i, movie := range movies {
		fmt.Println(movie, id)
		if movie.ID == id {
			fmt.Println("Found movie", movie)
			return movie, i, nil
		}
	}

	return Movie{}, 0, errors.New("invalid id")
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
