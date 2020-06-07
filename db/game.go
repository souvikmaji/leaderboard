package db

import "github.com/souvikmaji/leaderboard/models"

// SaveGame saves a new game in the database
func (db *DB) SaveGame(game *models.Game) (err error) {
	if err = db.DB.Create(game).Error; err != nil {
		return
	}

	return
}

// AllGames fetches all fantassy games from the database
// func (db *DB) AllGames(length, offset int64, orderBy, dir string) (games []*Game, recordsTotal, recordsFiltered int64, err error) {
func (db *DB) AllGames(length, offset int64) (games []*models.Game, recordsTotal, recordsFiltered int64, err error) {
	db.Model(&models.Game{}).Count(&recordsTotal)

	sqlDB := db.DB

	sqlDB = sqlDB.Model(&models.Game{}).Order("total_score desc")

	sqlDB.Count(&recordsFiltered)
	if err = sqlDB.Select("*, RANK () OVER (ORDER BY total_score desc) rank").Offset(offset).Limit(length).Find(&games).Error; err != nil {
		return
	}

	return
}
