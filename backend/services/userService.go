package services

import (
	"backend/models"
	"backend/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRequest struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func CreateUser(db *gorm.DB, request UserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: request.Username,
		FullName: request.FullName,
		Password: string(hashedPassword),
		Email:    request.Email,
		Role:     string("admin"),
	}

	return repositories.AddUser(db, user)
}

func EditUser(db *gorm.DB, userID uint, request UserRequest) error {
	var hashedPassword string
	if request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		hashedPassword = string(password)
	}

	updatedUser := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Username: request.Username,
		Password: hashedPassword, // Will update only if provided
	}

	return repositories.EditUser(db, userID, updatedUser)
}

func DeleteUser(db *gorm.DB, userID uint) error {
	return repositories.DeleteUser(db, userID)
}

func ChangeUserRole(db *gorm.DB, userID uint, newRole string) error {
	return repositories.ChangeUserRole(db, userID, newRole)
}

func GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	return repositories.GetUserByUsername(db, username)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT generates a JWT token for the authenticated user
func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"role":       role,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expires after 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
