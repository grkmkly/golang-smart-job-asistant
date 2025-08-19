package services

import (
	"errors"
	"smartjob/internal/auth"
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	DB           *gorm.DB
	UserService  *UserService
	TokenService *TokenService
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthService(db *gorm.DB, us *UserService, ts *TokenService) *AuthService {
	return &AuthService{db, us, ts}
}

func (s *AuthService) RegisterUser(req requests.RegisterRequest) (models.User, error) {

	user, err := mappers.RegisterRequestToUser(req)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *AuthService) LoginUser(req requests.LoginRequest) (*LoginResponse, error) {
	var user models.User

	result := s.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
	}

	if !utils.VerifyPassword(user.PasswordHash, req.Password) {
		return nil, errors.New("PasswordIsNotTrue")
	}

	//Access Token
	accessToken, err := auth.GenerateToken(user.ID, user.RoleID)
	if err != nil {
		return nil, err
	}

	//Refresh Token Create
	refreshTokenString, err := s.TokenService.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, errors.New("RefreshToken")
	}
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
	}, nil
}
