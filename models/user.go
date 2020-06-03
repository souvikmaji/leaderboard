package models

import "github.com/jinzhu/gorm"

// User reperesents a player
type User struct {
	gorm.Model
	Username string
	Email    string
}
