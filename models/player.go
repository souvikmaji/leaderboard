package models

// Player is a team member
type Player struct {
	id               int64
	name             string
	playingStyleDesc string `validate:"oneof=bowler wicketkeeper green batsman all-rounder"`
	score            int64  // 95.50 is represented as 95500
}

