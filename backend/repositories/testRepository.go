package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

// Addtest inserts a new test into a test
func AddTest(db *gorm.DB, test models.Test) error {
	return db.Create(&test).Error
}

// GetAllCourses retrieves all courses from the database
func GetAllTests(db *gorm.DB) ([]models.Test, error) {
	var tests []models.Test
	err := db.Find(&tests).Error
	if err != nil {
		return nil, err
	}
	return tests, nil
}

// GettestByID retrieves a test by its ID
func GetTestByID(db *gorm.DB, testID uint) (*models.Test, error) {
	var test models.Test
	err := db.Preload("QuestionOptions.AnswerOpts").Preload("QuestionEssays").First(&test, testID).Error
	return &test, err
}

// Updatetest updates a test's information
func UpdateTest(db *gorm.DB, testID uint, updatedtest models.Test) error {
	var test models.Test
	if err := db.First(&test, testID).Error; err != nil {
		return err
	}

	test.TestName = updatedtest.TestName
	test.Duration = updatedtest.Duration

	return db.Save(&test).Error
}

// Deletetest deletes a test by its ID
func DeleteTest(db *gorm.DB, testID uint) error {
	return db.Delete(&models.Test{}, testID).Error
}
