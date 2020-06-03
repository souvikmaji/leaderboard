package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/souvikmaji/leaderboard/models"
)

// Datastore interface defines operations possible on the database for this app
type Datastore interface {
	SaveTeam(team *models.Team) (err error)
	AllTeams(length, offset int64) (teams []*models.Team, recordsTotal, recordsFiltered int64, err error)
}

// DB is the wrapper for gorm db object
type DB struct {
	*gorm.DB
}

// NewDB creates a new database connection
func NewDB(dataSourceName string) (*DB, error) {
	// open new daatabase connection
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// check if connection is properly established
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}

	log.Println("Connection established with db")

	// return db wrapper
	return &DB{db}, nil
}
