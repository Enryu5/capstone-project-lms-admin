package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

// AddQuestionOpt inserts a new QuestionOpt into a course
func AddQuestionOpt(db *gorm.DB, questionOpt models.QuestionOpt) error {
	return db.Create(&questionOpt).Error
}

// GetAllQuestionOpt retrieves all QuestionOpt from the database
func GetAllQuestionOpt(db *gorm.DB) ([]models.QuestionOpt, error) {
	var questionOpts []models.QuestionOpt
	err := db.Find(&questionOpts).Error
	if err != nil {
		return nil, err
	}
	return questionOpts, nil
}

// GetQuestionOptByID retrieves a QuestionOpt by its ID
func GetQuestionOptByID(db *gorm.DB, questionOptID uint) (*models.QuestionOpt, error) {
	var questionOpt models.QuestionOpt
	err := db.Preload("AnswerOpt").First(&questionOpt, questionOptID).Error
	return &questionOpt, err
}

// UpdateQuestionOpt updates a QuestionOpt's information
func UpdateQuestionOpt(db *gorm.DB, questionOptID uint, updatedQuestionOpt models.QuestionOpt) error {
	var questionOpt models.QuestionOpt
	if err := db.First(&questionOpt, questionOptID).Error; err != nil {
		return err
	}

	questionOpt.Question = updatedQuestionOpt.Question

	return db.Save(&questionOpt).Error
}

// DeleteQuestionOpt deletes a QuestionOpt by its ID
func DeleteQuestionOpt(db *gorm.DB, questionOptID uint) error {
	return db.Delete(&models.QuestionOpt{}, questionOptID).Error
}
