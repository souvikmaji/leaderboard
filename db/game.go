package db

import "github.com/souvikmaji/leaderboard/models"

// SaveGame saves a new game in the database
func (db *DB) SaveGame(game *models.Game) (err error) {
	if err = db.DB.Create(game).Error; err != nil {
		return
	}

	return
}
