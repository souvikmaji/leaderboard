package models

import "github.com/jinzhu/gorm"

// Game is a enitity where the user participates for playing
type Game struct {
	gorm.Model
	GameID uint `gorm:"auto_increment" schema:"-"`
	Name   string
	UserID int64
	Score  float64
	Rank   int64 `gorm:"-" schema:"-"`
}
