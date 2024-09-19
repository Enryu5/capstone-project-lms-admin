package models

import (
	"gorm.io/gorm"
)

type QuestionEssay struct {
	gorm.Model
	Question string `gorm:"type:text" json:"question"`
	Answer   string `gorm:"type:text" json:"answer"`
	TestID   uint   `json:"test_id"`
}
