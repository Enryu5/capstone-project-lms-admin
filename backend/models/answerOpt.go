package models

import (
	"gorm.io/gorm"
)

type AnswerOpt struct {
	gorm.Model
	Answer        string `gorm:"type:text" json:"answer"`
	IsTrue        uint   `json:"is_true"`
	QuestionOptID uint   `json:"questionOpt_id"`
}
