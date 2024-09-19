package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

// AddCourse inserts a new course into the database
func AddCourse(db *gorm.DB, course models.Course) error {
	return db.Create(&course).Error
}

// GetAllCourses retrieves all courses from the database
func GetAllCourses(db *gorm.DB) ([]models.Course, error) {
	var courses []models.Course
	err := db.Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

// GetCourseByID retrieves a course by its ID
func GetCourseByID(db *gorm.DB, courseID uint) (*models.Course, error) {
	var course models.Course
	err := db.Preload("Modules").First(&course, courseID).Error
	return &course, err
}

// UpdateCourse updates a course's information
func UpdateCourse(db *gorm.DB, courseID uint, updatedCourse models.Course) error {
	var course models.Course
	if err := db.First(&course, courseID).Error; err != nil {
		return err
	}

	course.CourseName = updatedCourse.CourseName
	course.Desc = updatedCourse.Desc

	return db.Save(&course).Error
}

// DeleteCourse deletes a course by its ID
func DeleteCourse(db *gorm.DB, courseID uint) error {
	return db.Delete(&models.Course{}, courseID).Error
}
