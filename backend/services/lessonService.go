package services

import (
	"backend/models"
	"backend/repositories"

	"gorm.io/gorm"
)

// CreateLesson creates a new Lesson
func CreateLesson(db *gorm.DB, lesson models.Lesson) error {
	return repositories.AddLesson(db, lesson)
}

// GetLesson retrieves a Lesson by its ID
func GetLesson(db *gorm.DB, lessonID uint) (*models.Lesson, error) {
	return repositories.GetLessonByID(db, lessonID)
}

// UpdateLesson updates a Lesson
func UpdateLesson(db *gorm.DB, lessonID uint, updatedLesson models.Lesson) error {
	return repositories.UpdateLesson(db, lessonID, updatedLesson)
}

// DeleteLesson deletes a Lesson
func DeleteLesson(db *gorm.DB, lessonID uint) error {
	return repositories.DeleteLesson(db, lessonID)
}
