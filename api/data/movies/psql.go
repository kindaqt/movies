package data

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

////////////////////////////////////////
// Persister
//////////////////////////////////////

type PsqlStore struct {
	ConfigString string
	DB           *gorm.DB
}

func NewPsqlStore() Store {
	return &PsqlStore{
		ConfigString: fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v", host, port, user, dbname, password),
	}
}

// TODO: move to config
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "my-password"
	dbname   = "movies_db"
)

////////////////////////////////////////
// DB
//////////////////////////////////////

// Connect to DB
func (p *PsqlStore) Connect() {
	// Open DB
	db, err := gorm.Open("postgres", p.ConfigString)
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
	return nil, nil
}

// UpdateWatched updates the watch value value
func (p *PsqlStore) UpdateWatched(id string, value bool) error {
	return nil
}
