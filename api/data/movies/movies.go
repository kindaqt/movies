package data

import "log"

type Store interface {
	GetMovies() ([]Movie, error)
	UpdateWatched(id string, value bool) error
}

type store struct {
	Store
}

func NewStore() Store {
	return &JsonStore{}
}

func (c *store) GetMovies() ([]Movie, error) {
	log.Println("Movies: GetMovies")
	return c.Store.GetMovies()
}

func (c *store) UpdateWatched(id string, value bool) error {
	log.Println("Movies: UpdateWatched")
	return c.Store.UpdateWatched(id, value)
}
