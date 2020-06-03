package models

import "github.com/jinzhu/gorm"

// Team represets a fanatasy team created by the user consisting of many players
type Team struct {
	gorm.Model
	TeamID     uint `gorm:"AUTO_INCREMENT" schema:"-"`
	TeamName   string
	UserID     int64
	MatchID    int64
	CaptainID  int64
	VCaptainID int64
	TotalScore float64
	Rank       int64 `gorm:"-" schema:"-"`
}
