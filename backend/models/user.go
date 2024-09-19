package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `gorm:"uniqueIndex"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
