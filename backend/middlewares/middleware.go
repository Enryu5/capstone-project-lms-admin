package middlewares

import (
	"backend/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Define context keys to avoid collisions
const (
	ContextUserID   = "userID"
	ContextUserRole = "userRole"
)

// AuthMiddleware validates the JWT and extracts user information for Gin
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		userID, okID := (*claims)["user_id"].(float64)
		userRole, okRole := (*claims)["role"].(string)
		if !okID || !okRole {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("userID", uint(userID))
		c.Set("userRole", userRole)

		c.Next()
	}
}

// AdminMiddleware ensures only admins can access certain routes
func AdminMiddleware(c *gin.Context) {
	userRole, exists := c.Get(ContextUserRole)
	if !exists || userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admins only"})
		c.Abort()
		return
	}
	c.Next()
}

// InstructorMiddleware ensures only instructors and admins can access certain routes
func InstructorMiddleware(c *gin.Context) {
	userRole, exists := c.Get(ContextUserRole)
	if !exists || (userRole != "admin" && userRole != "instructor") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Instructors and Admins only"})
		c.Abort()
		return
	}
	c.Next()
}
