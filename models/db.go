package models

import (
	"github.com/jinzhu/gorm"
)

// Datastore ...
type Datastore interface {
	AllTeams(int64) ([]*Team, int64, error)
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
