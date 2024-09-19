package controllers

import (
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionEssayController struct {
	DB *gorm.DB
}

// CreateQuestionEssayHandler creates a new QuestionEssay
func CreateQuestionEssayHandler(c *gin.Context) {
	var questionEssay models.QuestionEssay
	if err := c.ShouldBindJSON(&questionEssay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)

	if err := repositories.AddQuestionEssay(db, questionEssay); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create QuestionEssay"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "QuestionEssay created successfully"})
}

func GetAllQuestionEssaysHandler(c *gin.Context) {
	db := c.MustGet("dbConn").(*gorm.DB)

	questionEssays, err := repositories.GetAllQuestionEssay(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve Question Essayions"})
		return
	}

	c.JSON(http.StatusOK, questionEssays)
}

// GetQuestionEssayHandler retrieves a QuestionEssay by its ID
func GetQuestionEssayHandler(c *gin.Context) {
	questionEssayID, _ := strconv.ParseUint(c.Param("questionEssayID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	QuestionEssay, err := repositories.GetQuestionEssayByID(db, uint(questionEssayID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QuestionEssay not found"})
		return
	}

	c.JSON(http.StatusOK, QuestionEssay)
}

// UpdateQuestionEssayHandler updates a QuestionEssay's details
func UpdateQuestionEssayHandler(c *gin.Context) {
	questionEssayID, _ := strconv.ParseUint(c.Param("questionEssayID"), 10, 64)
	var updatedQuestionEssay models.QuestionEssay
	if err := c.ShouldBindJSON(&updatedQuestionEssay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.UpdateQuestionEssay(db, uint(questionEssayID), updatedQuestionEssay); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update QuestionEssay"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "QuestionEssay updated successfully"})
}

// DeleteQuestionEssayHandler deletes a QuestionEssay
func DeleteQuestionEssayHandler(c *gin.Context) {
	questionEssayID, _ := strconv.ParseUint(c.Param("questionEssayID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.DeleteQuestionEssay(db, uint(questionEssayID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete QuestionEssay"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "QuestionEssay deleted successfully"})
}
