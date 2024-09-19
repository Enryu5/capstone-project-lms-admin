package controllers

import (
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnswerOptController struct {
	DB *gorm.DB
}

// CreateAnswerOptHandler creates a new AnswerOpt
func CreateAnswerOptHandler(c *gin.Context) {
	var AnswerOpt models.AnswerOpt
	if err := c.ShouldBindJSON(&AnswerOpt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)

	if err := repositories.AddAnswerOpt(db, AnswerOpt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create AnswerOpt"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "AnswerOpt created successfully"})
}

// GetAnswerOptHandler retrieves a AnswerOpt by its ID
func GetAnswerOptHandler(c *gin.Context) {
	AnswerOptID, _ := strconv.ParseUint(c.Param("AnswerOptID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	AnswerOpt, err := repositories.GetAnswerOptByID(db, uint(AnswerOptID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AnswerOpt not found"})
		return
	}

	c.JSON(http.StatusOK, AnswerOpt)
}

// UpdateAnswerOptHandler updates a AnswerOpt's details
func UpdateAnswerOptHandler(c *gin.Context) {
	AnswerOptID, _ := strconv.ParseUint(c.Param("AnswerOptID"), 10, 64)
	var updatedAnswerOpt models.AnswerOpt
	if err := c.ShouldBindJSON(&updatedAnswerOpt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.UpdateAnswerOpt(db, uint(AnswerOptID), updatedAnswerOpt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update AnswerOpt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "AnswerOpt updated successfully"})
}

// DeleteAnswerOptHandler deletes a AnswerOpt
func DeleteAnswerOptHandler(c *gin.Context) {
	AnswerOptID, _ := strconv.ParseUint(c.Param("AnswerOptID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.DeleteAnswerOpt(db, uint(AnswerOptID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete AnswerOpt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "AnswerOpt deleted successfully"})
}
