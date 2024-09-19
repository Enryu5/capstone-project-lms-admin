package services

import (
	"backend/repositories"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	return repositories.ValidateToken(s.db, tokenString)
}

func (s *AuthService) IsTokenBlacklisted(tokenString string) bool {
	return repositories.IsTokenBlacklisted(s.db, tokenString)
}
