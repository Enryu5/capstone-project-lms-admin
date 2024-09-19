package controllers

import (
	"backend/repositories"
	"backend/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

// CreateUser handles the creation of a new user
func (uc *AuthController) RegisterHandler(c *gin.Context) {
	var request services.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := services.CreateUser(uc.DB, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// LoginHandler authenticates the user and generates a JWT token
func (uc *AuthController) LoginHandler(c *gin.Context) {
	var loginRequest services.UserRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := c.MustGet("dbConn").(*gorm.DB)

	// Check if user exists
	user, err := services.GetUserByUsername(db, loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verify password
	if err := services.VerifyPassword(user.Password, loginRequest.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := services.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// LogoutHandler invalidates the user's JWT token
func (uc *AuthController) LogoutHandler(c *gin.Context) {
	db := c.MustGet("dbConn").(*gorm.DB) // Use the correct "dbConn" key
	tokenString := extractToken(c)

	// Blacklist token to invalidate it
	err := repositories.BlacklistToken(db, tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to blacklist token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Helper to extract the token from the Authorization header
func extractToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer ") {
		return strings.TrimPrefix(bearerToken, "Bearer ")
	}
	return ""
}
