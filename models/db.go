package models

import (
	"github.com/jinzhu/gorm"
)

// Datastore ...
type Datastore interface {
	AllTeams(int64) ([]*Team, int64, error)
}

type DB struct {
	*gorm.DB
}

// NewDB creates a new database connection
func NewDB(dataSourceName string) (*DB, error) {
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
