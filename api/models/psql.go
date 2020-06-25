package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// TODO: move to config
const (
	driver   = "postgres"
	host     = "db"
	port     = 5432
	user     = "banana"
	password = "banana"
	dbname   = "banana"
)

////////////////////////////////////////
// Persister
//////////////////////////////////////

type PsqlStore struct {
	ConfigString string
	DB           *gorm.DB
}

func NewPsqlStore() Store {
	s := &PsqlStore{
		ConfigString: fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, dbname, password),
	}
	s.Connect()

	return s
}

////////////////////////////////////////
// DB
//////////////////////////////////////

// Connect to DB
func (p *PsqlStore) Connect() {
	// Open DB
	db, err := gorm.Open(driver, p.ConfigString)
	if err != nil {
		log.Panic(err)
	}

	if err := db.DB().Ping(); err != nil {
		log.Panic(err)
	}

	// Attach DB to persister
	p.DB = db
}

func (p *PsqlStore) Close() error {
	if err := p.DB.Close(); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////
// Methods
//////////////////////////////////////

// GetMovies() returns all movies
func (p *PsqlStore) GetMovies() ([]Movie, error) {
	log.Println("Psql: GetMovies")
	var movies []Movie
	err := p.DB.Find(&movies).Error
	return movies, err
}

// UpdateWatched() updates the watch value value
func (p *PsqlStore) UpdateWatched(id string, value bool) error {
	log.Println("Psql: UpdateWatched")
	err := p.DB.Model(&Movie{ID: id}).Update("watched", value).Error
	return err
}

// DeleteMovie() deletes a movie from the database
func (p *PsqlStore) DeleteMovie(id string) error {
	log.Println("Psql: DeleteMovie")
	err := p.DB.Delete(&Movie{ID: id}).Error
	return err
}

// CreateMovie() creates a movie record in the database
func (p *PsqlStore) CreateMovie(movie *Movie) error {
	log.Println("Psql: CreateMovie")
	// TODO:
	// - title must be unique
	err := p.DB.Create(movie).Error
	return err
}
