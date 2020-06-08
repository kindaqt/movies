package model

import (
	"log"
)

type Store interface {
	GetMovies() ([]Movie, error)
	UpdateWatched(id string, value bool) error
	DeleteMovie(id string) error
	CreateMovie(movie *Movie) error
}

type store struct {
	Store
}

func NewStore(name string) Store {
	var myStore Store
	switch name {
	case "json":
		var err error
		myStore, err = NewJsonStore("./data/movies/movies.json")
		if err != nil {
			panic(err)
		}
	case "psql":
		myStore = NewPsqlStore()
	default:
		panic("invalid store name of " + name + "please use a valid store name and try again.")
	}

	return myStore
}

func (c *store) GetMovies() ([]Movie, error) {
	log.Println("Movies: GetMovies")
	return c.Store.GetMovies()
}

func (c *store) UpdateWatched(id string, value bool) error {
	log.Println("Movies: UpdateWatched")
	return c.Store.UpdateWatched(id, value)
}

func (c *store) DeleteMovie(id string) error {
	log.Println("Movies: DeleteMovie")
	return c.Store.DeleteMovie(id)
}

func (c *store) CreateMovie(movie *Movie) error {
	log.Println("Movies: CreateMovie")
	return c.Store.CreateMovie(movie)
}
