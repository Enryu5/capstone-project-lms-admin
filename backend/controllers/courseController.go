package controllers

import (
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CourseController struct {
	DB *gorm.DB
}

// CreateCourseHandler creates a new course
func CreateCourseHandler(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)

	if err := repositories.AddCourse(db, course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create course"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
}

func GetAllCoursesHandler(c *gin.Context) {
	db := c.MustGet("dbConn").(*gorm.DB)

	courses, err := repositories.GetAllCourses(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// GetCourseHandler retrieves a course by its ID
func GetCourseByIDHandler(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("courseID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	course, err := repositories.GetCourseByID(db, uint(courseID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// UpdateCourseHandler updates a course's details
func UpdateCourseHandler(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("courseID"), 10, 64)
	var updatedCourse models.Course
	if err := c.ShouldBindJSON(&updatedCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.UpdateCourse(db, uint(courseID), updatedCourse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

// DeleteCourseHandler deletes a course
func DeleteCourseHandler(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("courseID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.DeleteCourse(db, uint(courseID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
