package models

import (
	"gorm.io/gorm"
)

type Lesson struct {
	gorm.Model
	LessonName string `gorm:"uniqueIndex" json:"title"`
	Content    string `gorm:"type:text" json:"content"`
	ModuleID   uint   `json:"module_id"`
}
