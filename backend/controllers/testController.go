package controllers

import (
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TestController struct {
	DB *gorm.DB
}

// CreateTestHandler creates a new Test
func CreateTestHandler(c *gin.Context) {
	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)

	if err := repositories.AddTest(db, test); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Test"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Test created successfully"})
}

// GetTestHandler retrieves a Test by its ID
func GetTestHandler(c *gin.Context) {
	testID, _ := strconv.ParseUint(c.Param("testID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	Test, err := repositories.GetTestByID(db, uint(testID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test not found"})
		return
	}

	c.JSON(http.StatusOK, Test)
}

// UpdateTestHandler updates a Test's details
func UpdateTestHandler(c *gin.Context) {
	testID, _ := strconv.ParseUint(c.Param("testID"), 10, 64)
	var updatedTest models.Test
	if err := c.ShouldBindJSON(&updatedTest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.UpdateTest(db, uint(testID), updatedTest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Test"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test updated successfully"})
}

// DeleteTestHandler deletes a Test
func DeleteTestHandler(c *gin.Context) {
	testID, _ := strconv.ParseUint(c.Param("testID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.DeleteTest(db, uint(testID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Test"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Test deleted successfully"})
}
