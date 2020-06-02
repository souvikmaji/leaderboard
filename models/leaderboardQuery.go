package models

// LeaderboardQuery holds the query parameters for getting the leaderboard
type LeaderboardQuery struct {
	Draw   int64
	Length int64
	Start  int64
}
