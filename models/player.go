package models

// Player is a team member
type Player struct {
	ID               uint `gorm:"primary_key"`
	Name             string
	PlayingStyleDesc string  `validate:"oneof=bowler wicketkeeper green batsman all-rounder"`
	Score            float64 // 95.50 is represented as 95500
}
