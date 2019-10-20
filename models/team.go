package models

// Team represets a fanatasy team created by the user consisting of many players
type Team struct {
	teamID     int64
	teamName   string
	userID     int64
	matchID    int64
	captainID  int64
	vcaptainID int64
	totalScore int64
	players    []*Player
}
