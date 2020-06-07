package db

import (
	"github.com/jinzhu/gorm"
	"github.com/souvikmaji/leaderboard/models"
)

// Datastore interface defines operations possible on the database for this app
type Datastore interface {
	// User operations
	SaveUser(user *models.User) (err error)
	GetUser(userQuery *models.User) (user *models.User, err error)

	// Game operations
	SaveGame(game *models.Game) (err error)
	AllGames(length, offset int64) (games []*models.Game, recordsTotal, recordsFiltered int64, err error)
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

	// return db wrapper
	return &DB{db}, nil
}
