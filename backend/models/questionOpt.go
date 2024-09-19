package models

import (
	"gorm.io/gorm"
)

type QuestionOpt struct {
	gorm.Model
	Question   string      `gorm:"type:text" json:"question"`
	AnswerOpts []AnswerOpt `gorm:"foreignKey:QuestionOptID" json:"answers"`
	TestID     uint        `json:"test_id"`
}
