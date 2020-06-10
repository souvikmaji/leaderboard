package models

// GameUser represents each games an user has joined
type GameUser struct {
	GameID uint
	UserID uint
	Score  float64
	Rank   int64 `gorm:"-" schema:"-"`
	Game   *Game
}
