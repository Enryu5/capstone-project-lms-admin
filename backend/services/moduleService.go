package services

import (
	"backend/models"
	"backend/repositories"

	"gorm.io/gorm"
)

// CreateModule creates a new module
func CreateModule(db *gorm.DB, module models.Module) error {
	return repositories.AddModule(db, module)
}

// GetModule retrieves a Module by its ID
func GetModule(db *gorm.DB, moduleID uint) (*models.Module, error) {
	return repositories.GetModuleByID(db, moduleID)
}

// UpdateModule updates a Module
func UpdateModule(db *gorm.DB, moduleID uint, updatedModule models.Module) error {
	return repositories.UpdateModule(db, moduleID, updatedModule)
}

// DeleteModule deletes a Module
func DeleteModule(db *gorm.DB, moduleID uint) error {
	return repositories.DeleteModule(db, moduleID)
}
