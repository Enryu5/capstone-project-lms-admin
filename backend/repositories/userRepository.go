package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

// AddUser inserts a new user into the database
func AddUser(db *gorm.DB, user models.User) error {
	return db.Create(&user).Error
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// DeleteUser deletes a user by their ID
func DeleteUser(db *gorm.DB, userID uint) error {
	return db.Delete(&models.User{}, userID).Error
}

// EditUser updates the user's information (except the role)
func EditUser(db *gorm.DB, userID uint, updatedUser models.User) error {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	// Update only specific fields
	user.FullName = updatedUser.FullName
	user.Email = updatedUser.Email
	user.Username = updatedUser.Username

	if updatedUser.Password != "" {
		user.Password = updatedUser.Password // Password should be hashed before being passed here
	}

	return db.Save(&user).Error
}

// ChangeUserRole updates the user's role (e.g., admin, student)
func ChangeUserRole(db *gorm.DB, userID uint, newRole string) error {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	user.Role = newRole
	return db.Save(&user).Error
}
