package services

import (
	"backend/models"
	"backend/repositories"

	"gorm.io/gorm"
)

// CreateCourse creates a new course
func CreateCourse(db *gorm.DB, course models.Course) error {
	return repositories.AddCourse(db, course)
}

// GetCourse retrieves a course by its ID
func GetCourse(db *gorm.DB, courseID uint) (*models.Course, error) {
	return repositories.GetCourseByID(db, courseID)
}

// UpdateCourse updates a course
func UpdateCourse(db *gorm.DB, courseID uint, updatedCourse models.Course) error {
	return repositories.UpdateCourse(db, courseID, updatedCourse)
}

// DeleteCourse deletes a course
func DeleteCourse(db *gorm.DB, courseID uint) error {
	return repositories.DeleteCourse(db, courseID)
}
