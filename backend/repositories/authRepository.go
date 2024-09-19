package repositories

import (
	"backend/models"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// BlacklistToken blacklists the given token by adding it to the database
func BlacklistToken(db *gorm.DB, tokenString string) error {
	blacklistEntry := models.Blacklist{Token: tokenString}
	return db.Create(&blacklistEntry).Error
}

// IsTokenBlacklisted checks if the given token is blacklisted
func IsTokenBlacklisted(db *gorm.DB, tokenString string) bool {
	var blacklist models.Blacklist
	err := db.Where("token = ?", tokenString).First(&blacklist).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// ValidateToken checks the validity of a given JWT token and returns the claims
func ValidateToken(db *gorm.DB, tokenString string) (*jwt.MapClaims, error) {
	// Check if the token is blacklisted
	if IsTokenBlacklisted(db, tokenString) {
		return nil, errors.New("token has been invalidated")
	}

	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// RefreshToken generates a new token for the user using existing claims
func RefreshToken(db *gorm.DB, oldTokenString string) (string, error) {
	claims, err := ValidateToken(db, oldTokenString)
	if err != nil {
		return "", err
	}

	// Check if token is expired (this is optional based on your needs)
	exp, ok := (*claims)["exp"].(float64)
	if !ok || time.Now().Unix() > int64(exp) {
		return "", errors.New("token has expired")
	}

	// Generate a new token using the existing claims (minus the old expiration time)
	newClaims := jwt.MapClaims{
		"authorized": (*claims)["authorized"],
		"user_id":    (*claims)["user_id"],
		"role":       (*claims)["role"],
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expires after 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	return token.SignedString(jwtSecret)
}
