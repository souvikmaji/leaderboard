package models

import "github.com/jinzhu/gorm"

// Game is a enitity where the user participates for playing
type Game struct {
	gorm.Model `schema:"-"`
	Name       string
}
