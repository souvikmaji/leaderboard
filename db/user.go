package db

import "github.com/souvikmaji/leaderboard/models"

// SaveUser saves a new user in the database
func (db *DB) SaveUser(user *models.User) (err error) {
	if err = db.DB.Create(user).Error; err != nil {
		return
	}

	return
}

// GetUser searches user by user id
func (db *DB) GetUser(user *models.User) (*models.User, error) {
	if err := db.DB.Where(user).Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
