package models

import (
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	TestName       string          `gorm:"uniqueIndex" json:"title"`
	Duration       int             `json:"duration"`
	QuestionOpts   []QuestionOpt   `gorm:"foreignKey:TestID" json:"questionOpts"`
	QuestionEssays []QuestionEssay `gorm:"foreignKey:TestID" json:"questionEssays"`
	ModuleID       uint            `json:"module_id"`
}
