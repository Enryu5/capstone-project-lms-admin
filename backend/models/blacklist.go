package models

import "gorm.io/gorm"

// Blacklist represents the tokens that are no longer valid
type Blacklist struct {
	gorm.Model
	Token string `gorm:"unique"`
}
