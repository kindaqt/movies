package data

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
		ConfigString: fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v", host, port, user, dbname, password),
	}
	s.Connect()

	return s
}

// TODO: move to config
const (
	driver   = "postgres"
	host     = "localhost"
	port     = 5432
	user     = "docker"
	password = "docker"
	dbname   = "docker"
)

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

// GetMovies returns all movies
func (p *PsqlStore) GetMovies() ([]Movie, error) {
	var movies []Movie
	err := p.DB.Find(&movies).Error
	return movies, err
}

// UpdateWatched updates the watch value value
func (p *PsqlStore) UpdateWatched(id string, value bool) error {
	return nil
}

// DeleteMovie deletes a movie from the database
func (p *PsqlStore) DeleteMovie(id string) error {
	log.Println("Psql: DeleteMovie")
	err := p.DB.Delete(&Movie{ID: id}).Error
	return err
}
