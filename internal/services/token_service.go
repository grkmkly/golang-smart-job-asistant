package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"smartjob/internal/auth"
	"smartjob/internal/models"
	"time"

	"gorm.io/gorm"
)

type TokenService struct {
	DB          *gorm.DB
	UserService *UserService
}

func NewTokenService(db *gorm.DB, us *UserService) *TokenService {
	return &TokenService{
		db,
		us,
	}
}
func GenerateSecureRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (s *TokenService) CreateRefreshToken(userID uint) (string, error) {
	refreshTokenString, err := GenerateSecureRandomToken(32)
	if err != nil {
		return "", errors.New("RefreshTokenError")
	}
	refreshTokenHash := HashToken(refreshTokenString)
	refreshTokenExpiresAt := time.Now().Add(7 * 24 * time.Hour)

	newRefreshToken := models.RefreshToken{
		UserID:    userID,
		TokenHash: refreshTokenHash,
		ExpiresAt: refreshTokenExpiresAt,
	}

	if err := s.DB.Create(&newRefreshToken).Error; err != nil {
		return "", errors.New("RefreshTokennotsave")
	}

	return refreshTokenString, nil
}

func (s *TokenService) RefreshAccessToken(refreshTokenString string) (string, error) {
	refreshTokenHash := HashToken(refreshTokenString)

	var refreshToken models.RefreshToken
	result := s.DB.Where("token_hash = ?", refreshTokenHash).First(&refreshToken)
	if result.Error != nil {
		return "", errors.New("geçersiz veya bulunamayan refresh token")
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		s.DB.Delete(&refreshToken)
		return "", errors.New("süresi dolmuş refresh token, lütfen tekrar giriş yapın")
	}

	roleID, err := s.UserService.GetUserRoleID(refreshToken.UserID)
	if err != nil {
		return "", err
	}

	newAccessToken, err := auth.GenerateToken(refreshToken.UserID, roleID)
	if err != nil {
		return "", errors.New("yeni access token oluşturulamadı")
	}

	return newAccessToken, nil
}

func (s *TokenService) Validate(tokenString string) (*models.User, error) {
	var refreshToken models.RefreshToken

	err := s.DB.Preload("User").Where("token = ?", tokenString).First(&refreshToken).Error
	if err != nil {
		return nil, errors.New("geçersiz refresh token")
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		s.DB.Delete(&refreshToken)
		return nil, errors.New("süresi dolmuş refresh token, lütfen tekrar giriş yapın")
	}

	return &refreshToken.User, nil
}
