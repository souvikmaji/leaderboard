package models

// Team represets a fanatasy team created by the user consisting of many players
type Team struct {
	TeamID     uint `gorm:"primary_key"`
	TeamName   string
	UserID     int64
	MatchID    int64
	CaptainID  int64
	VCaptainID int64
	TotalScore float64
	Rank       int64     `gorn:"-"`
	Players    []*Player `gorm:"many2many:team_players;association_foreignkey:id;foreignkey:team_id"`
}

// AllTeams fetches all fantassy teams from the database
// func (db *DB) AllTeams(length, offset int64, orderBy, dir string) (teams []*Team, recordsTotal, recordsFiltered int64, err error) {
func (db *DB) AllTeams(length, offset int64) (teams []*Team, recordsTotal, recordsFiltered int64, err error) {
	db.Model(&Team{}).Count(&recordsTotal)

	sqlDB := db.DB

	sqlDB = sqlDB.Model(&Team{}).Order("total_score desc")

	sqlDB.Count(&recordsFiltered)
	sqlDB.Select("*, RANK () OVER ( ORDER BY total_score desc) rank").Offset(offset).Limit(length).Find(&teams)

	return
}
