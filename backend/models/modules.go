package models

import (
	"gorm.io/gorm"
)

type Module struct {
	gorm.Model
	ModuleName string   `gorm:"uniqueIndex" json:"title"`
	Desc       string   `gorm:"text" json:"description"`
	Tests      []Test   `gorm:"foreignKey:ModuleID" json:"tests"`
	Lessons    []Lesson `gorm:"foreignKey:ModuleID" json:"lessons"`
	CourseID   uint     `json:"course_id"`
}
