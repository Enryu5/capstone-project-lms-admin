package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

// AddAnswerOpt inserts a new answerOpt into a module
func AddAnswerOpt(db *gorm.DB, answerOpt models.AnswerOpt) error {
	return db.Create(&answerOpt).Error
}

// GetanswerOptByID retrieves a answerOpt by its ID
func GetAnswerOptByID(db *gorm.DB, answerOptID uint) (*models.AnswerOpt, error) {
	var answerOpt models.AnswerOpt
	err := db.First(&answerOpt, answerOptID).Error
	return &answerOpt, err
}

// UpdateanswerOpt updates a answerOpt's information
func UpdateAnswerOpt(db *gorm.DB, answerOptID uint, updatedAnswerOpt models.AnswerOpt) error {
	var answerOpt models.AnswerOpt
	if err := db.First(&answerOpt, answerOptID).Error; err != nil {
		return err
	}

	answerOpt.Answer = updatedAnswerOpt.Answer
	answerOpt.IsTrue = updatedAnswerOpt.IsTrue

	return db.Save(&answerOpt).Error
}

// DeleteanswerOpt deletes a answerOpt by its ID
func DeleteAnswerOpt(db *gorm.DB, answerOptID uint) error {
	return db.Delete(&models.AnswerOpt{}, answerOptID).Error
}
