package db

import (
	"fmt"

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

	// Game to User operations
	SaveGameUser(game *models.GameUser) (err error)
	GetAllSortedGameUser(length, offset int64) (games []*models.GameUser, recordsTotal, recordsFiltered int64, err error)
}

// DB is the wrapper for gorm db object
type DB struct {
	*gorm.DB
}

// NewDB creates a new database connection
func NewDB(c *models.Configurations) (*DB, error) {
	// open new daatabase connection
	db, err := gorm.Open("postgres", getDbURI(c))
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

// getDbURI parses configurations and returns database server address in
// postgreSQL connection string format
func getDbURI(c *models.Configurations) string {
	if c.Database.URL != "" {
		return c.Database.URL
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Name)

	if c.Database.Password != "" {
		connectionString = fmt.Sprintf("%s password=%s", connectionString, c.Database.Password)
	}

	return connectionString
}
