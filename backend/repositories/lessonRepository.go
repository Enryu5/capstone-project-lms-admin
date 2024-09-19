package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

// AddLesson inserts a new lesson into a module
func AddLesson(db *gorm.DB, lesson models.Lesson) error {
	return db.Create(&lesson).Error
}

// GetAllCourses retrieves all courses from the database
func GetAllLessons(db *gorm.DB) ([]models.Lesson, error) {
	var lessons []models.Lesson
	err := db.Find(&lessons).Error
	if err != nil {
		return nil, err
	}
	return lessons, nil
}

// GetLessonByID retrieves a lesson by its ID
func GetLessonByID(db *gorm.DB, lessonID uint) (*models.Lesson, error) {
	var lesson models.Lesson
	err := db.First(&lesson, lessonID).Error
	return &lesson, err
}

// UpdateLesson updates a lesson's information
func UpdateLesson(db *gorm.DB, lessonID uint, updatedLesson models.Lesson) error {
	var lesson models.Lesson
	if err := db.First(&lesson, lessonID).Error; err != nil {
		return err
	}

	lesson.LessonName = updatedLesson.LessonName
	lesson.Content = updatedLesson.Content

	return db.Save(&lesson).Error
}

// DeleteLesson deletes a lesson by its ID
func DeleteLesson(db *gorm.DB, lessonID uint) error {
	return db.Delete(&models.Lesson{}, lessonID).Error
}
