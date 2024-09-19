package controllers

import (
	"backend/models"
	"backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ModuleController struct {
	DB *gorm.DB
}

// CreateModuleHandler creates a new Module
func CreateModuleHandler(c *gin.Context) {
	var module models.Module
	if err := c.ShouldBindJSON(&module); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)

	if err := repositories.AddModule(db, module); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create Module"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Module created successfully"})
}

// GetModuleHandler retrieves a Module by its ID
func GetModuleHandler(c *gin.Context) {
	moduleID, _ := strconv.ParseUint(c.Param("moduleID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	Module, err := repositories.GetModuleByID(db, uint(moduleID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
		return
	}

	c.JSON(http.StatusOK, Module)
}

// UpdateModuleHandler updates a Module's details
func UpdateModuleHandler(c *gin.Context) {
	moduleID, _ := strconv.ParseUint(c.Param("moduleID"), 10, 64)
	var updatedModule models.Module
	if err := c.ShouldBindJSON(&updatedModule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.UpdateModule(db, uint(moduleID), updatedModule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update Module"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Module updated successfully"})
}

// DeleteModuleHandler deletes a Module
func DeleteModuleHandler(c *gin.Context) {
	moduleID, _ := strconv.ParseUint(c.Param("moduleID"), 10, 64)

	db := c.MustGet("dbConn").(*gorm.DB)
	if err := repositories.DeleteModule(db, uint(moduleID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete Module"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Module deleted successfully"})
}
