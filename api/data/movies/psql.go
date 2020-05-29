package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

////////////////////////////////////////
// Persister
//////////////////////////////////////

type PsqlStore struct {
	ConfigString string
	DB           *gorm.DB
}

func NewPsqlStore() *PsqlStore {
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
// Methods
//////////////////////////////////////

// Connect to DB
func (p *PsqlStore) Connect() error {
	// Open DB
	db, err := gorm.Open("postgres", p.ConfigString)
	if err != nil {
		return err
	}

	// Attach DB to persister
	p.DB = db

	return nil
}
