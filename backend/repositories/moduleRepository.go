package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

// AddModule inserts a new module into a course
func AddModule(db *gorm.DB, module models.Module) error {
	return db.Create(&module).Error
}

// GetAllCourses retrieves all courses from the database
func GetAllModules(db *gorm.DB) ([]models.Module, error) {
	var modules []models.Module
	err := db.Find(&modules).Error
	if err != nil {
		return nil, err
	}
	return modules, nil
}

// GetModuleByID retrieves a module by its ID
func GetModuleByID(db *gorm.DB, moduleID uint) (*models.Module, error) {
	var module models.Module
	err := db.Preload("Lessons").Preload("Tests").First(&module, moduleID).Error
	return &module, err
}

// UpdateModule updates a module's information
func UpdateModule(db *gorm.DB, moduleID uint, updatedModule models.Module) error {
	var module models.Module
	if err := db.First(&module, moduleID).Error; err != nil {
		return err
	}

	module.ModuleName = updatedModule.ModuleName
	module.Desc = updatedModule.Desc

	return db.Save(&module).Error
}

// DeleteModule deletes a module by its ID
func DeleteModule(db *gorm.DB, moduleID uint) error {
	return db.Delete(&models.Module{}, moduleID).Error
}
