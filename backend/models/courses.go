package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseName string   `gorm:"uniqueIndex" json:"title"`
	Desc       string   `gorm:"type:text" json:"description"`
	CoursePic  string   `json:"picture"`
	Modules    []Module `gorm:"foreignKey:CourseID" json:"modules"`
}
