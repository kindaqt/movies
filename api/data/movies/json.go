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

func NewJsonStore(filePath string) (Store, error) {
	log.Println("Json: NewJsonStore")
	absPath, _ := filepath.Abs(filePath)
	if fileExists(absPath) {
		return &JsonStore{
			JsonFilePath: absPath,
		}, nil
	} else {
		return nil, errors.New("File does not exist.")
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

	if err := writeJSONFile(p.JsonFilePath, movies); err != nil {
		return err
	}

	return nil
}

func (p *JsonStore) DeleteMovie(id string) error {
	movies, err := p.GetMovies()
	if err != nil {
		return err
	}

	var newMovies []Movie
	var found bool = false
	for i := 0; i < len(movies); i++ {
		movie := movies[i]
		if movie.ID == id {
			found = true
			newMovies = append(newMovies, movies[i+1:]...)
		} else {
			newMovies = append(newMovies, movie)
		}
	}

	if !found {
		return errors.New("invalid id")
	}

	if err := writeJSONFile(p.JsonFilePath, newMovies); err != nil {
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

func writeJSONFile(filePath string, data interface{}) error {
	log.Println("Marshalling JSON...")
	file, err := json.Marshal(data)
	if err != nil {
		return err
	}
	log.Println("Writing JSON file...")
	if err := ioutil.WriteFile(filePath, file, 0644); err != nil {
		return err
	}
	return nil
}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
