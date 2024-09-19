package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

// AddQuestionEssay inserts a new QuestionEssay into a course
func AddQuestionEssay(db *gorm.DB, questionEssay models.QuestionEssay) error {
	return db.Create(&questionEssay).Error
}

// GetAllQuestionEssay retrieves all QuestionEssay from the database
func GetAllQuestionEssay(db *gorm.DB) ([]models.QuestionEssay, error) {
	var questionEssays []models.QuestionEssay
	err := db.Find(&questionEssays).Error
	if err != nil {
		return nil, err
	}
	return questionEssays, nil
}

// GetQuestionEssayByID retrieves a QuestionEssay by its ID
func GetQuestionEssayByID(db *gorm.DB, questionEssayID uint) (*models.QuestionEssay, error) {
	var questionEssay models.QuestionEssay
	err := db.First(&questionEssay, questionEssayID).Error
	return &questionEssay, err
}

// UpdateQuestionEssay updates a QuestionEssay's information
func UpdateQuestionEssay(db *gorm.DB, questionEssayID uint, updatedQuestionEssay models.QuestionEssay) error {
	var questionEssay models.QuestionEssay
	if err := db.First(&questionEssay, questionEssayID).Error; err != nil {
		return err
	}

	questionEssay.Question = updatedQuestionEssay.Question
	questionEssay.Answer = updatedQuestionEssay.Answer

	return db.Save(&questionEssay).Error
}

// DeleteQuestionEssay deletes a QuestionEssay by its ID
func DeleteQuestionEssay(db *gorm.DB, questionEssayID uint) error {
	return db.Delete(&models.QuestionEssay{}, questionEssayID).Error
}
