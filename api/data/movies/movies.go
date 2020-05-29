package data

type Store interface {
	GetMovies() (interface{}, error)
	UpdateWatched(id string, value bool) error
}
