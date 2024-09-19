package controllers

import (
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionOptController struct {
	DB *gorm.DB
}

// CreateQuestionOptHandler creates a new QuestionOpt
func CreateQuestionOptHandler(c *gin.Context) {
	var questionOpt models.QuestionOpt
	if err := c.ShouldBindJSON(&questionOpt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)

	if err := repositories.AddQuestionOpt(db, questionOpt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create QuestionOpt"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "QuestionOpt created successfully"})
}

func GetAllQuestionOptsHandler(c *gin.Context) {
	db := c.MustGet("dbConn").(*gorm.DB)

	questionOpts, err := repositories.GetAllQuestionOpt(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve Question Options"})
		return
	}

	c.JSON(http.StatusOK, questionOpts)
}

// GetQuestionOptHandler retrieves a QuestionOpt by its ID
func GetQuestionOptHandler(c *gin.Context) {
	questionOptID, _ := strconv.ParseUint(c.Param("questionOptID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	QuestionOpt, err := repositories.GetQuestionOptByID(db, uint(questionOptID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QuestionOpt not found"})
		return
	}

	c.JSON(http.StatusOK, QuestionOpt)
}

// UpdateQuestionOptHandler updates a QuestionOpt's details
func UpdateQuestionOptHandler(c *gin.Context) {
	questionOptID, _ := strconv.ParseUint(c.Param("questionOptID"), 10, 64)
	var updatedQuestionOpt models.QuestionOpt
	if err := c.ShouldBindJSON(&updatedQuestionOpt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.UpdateQuestionOpt(db, uint(questionOptID), updatedQuestionOpt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update QuestionOpt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "QuestionOpt updated successfully"})
}

// DeleteQuestionOptHandler deletes a QuestionOpt
func DeleteQuestionOptHandler(c *gin.Context) {
	questionOptID, _ := strconv.ParseUint(c.Param("questionOptID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.DeleteQuestionOpt(db, uint(questionOptID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete QuestionOpt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "QuestionOpt deleted successfully"})
}
