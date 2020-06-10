package db

import (
	"github.com/souvikmaji/leaderboard/models"
)

// SaveGameUser maps a user to a new game
func (db *DB) SaveGameUser(gameUser *models.GameUser) (err error) {
	if err = db.DB.Create(gameUser).Error; err != nil {
		return
	}

	return
}

// GetAllSortedGameUser fetches all user games based on user score
func (db *DB) GetAllSortedGameUser(length, offset int64) (gameUser []*models.GameUser, recordsTotal, recordsFiltered int64, err error) {
	db.Model(&models.GameUser{}).Count(&recordsTotal)

	sqlDB := db.DB

	sqlDB = sqlDB.Model(&models.GameUser{}).Preload("Game").Order("score desc")

	sqlDB.Count(&recordsFiltered)
	if err = sqlDB.Select("*, RANK () OVER (ORDER BY score desc) rank").Offset(offset).Limit(length).Find(&gameUser).Error; err != nil {
		return
	}

	return
}
