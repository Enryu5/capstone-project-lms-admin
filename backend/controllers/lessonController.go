package controllers

import (
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LessonController struct {
	DB *gorm.DB
}

// CreateLessonHandler creates a new Lesson
func CreateLessonHandler(c *gin.Context) {
	var lesson models.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)

	if err := repositories.AddLesson(db, lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Lesson"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Lesson created successfully"})
}

// GetLessonHandler retrieves a Lesson by its ID
func GetLessonHandler(c *gin.Context) {
	lessonID, _ := strconv.ParseUint(c.Param("lessonID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	Lesson, err := repositories.GetLessonByID(db, uint(lessonID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	c.JSON(http.StatusOK, Lesson)
}

// UpdateLessonHandler updates a Lesson's details
func UpdateLessonHandler(c *gin.Context) {
	lessonID, _ := strconv.ParseUint(c.Param("lessonID"), 10, 64)
	var updatedLesson models.Lesson
	if err := c.ShouldBindJSON(&updatedLesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.UpdateLesson(db, uint(lessonID), updatedLesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Lesson"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson updated successfully"})
}

// DeleteLessonHandler deletes a Lesson
func DeleteLessonHandler(c *gin.Context) {
	lessonID, _ := strconv.ParseUint(c.Param("lessonID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.DeleteLesson(db, uint(lessonID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Lesson"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted successfully"})
}
