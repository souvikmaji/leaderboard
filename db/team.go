package db

import "github.com/souvikmaji/leaderboard/models"

// SaveTeam saves a new team in the database
func (db *DB) SaveTeam(team *models.Team) (err error) {
	if err = db.DB.Create(team).Error; err != nil {
		return
	}

	return
}

// AllTeams fetches all fantassy teams from the database
// func (db *DB) AllTeams(length, offset int64, orderBy, dir string) (teams []*Team, recordsTotal, recordsFiltered int64, err error) {
func (db *DB) AllTeams(length, offset int64) (teams []*models.Team, recordsTotal, recordsFiltered int64, err error) {
	db.Model(&models.Team{}).Count(&recordsTotal)

	sqlDB := db.DB

	sqlDB = sqlDB.Model(&models.Team{}).Order("total_score desc")

	sqlDB.Count(&recordsFiltered)
	if err = sqlDB.Select("*, RANK () OVER (ORDER BY total_score desc) rank").Offset(offset).Limit(length).Find(&teams).Error; err != nil {
		return
	}

	return
}
