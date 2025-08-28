package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/responses"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db}
}
func (s *UserService) CreateUser(user *models.User) error {
	result := s.DB.Create(&user)
	return result.Error
}

func (s *UserService) GetUserProfile(userID uint) (*responses.UserResponse, error) {
	var user models.User
	result := s.DB.First(&user, userID)
	response := mappers.UserModelToResponse(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return response, nil
}
func (s *UserService) GetUserRoleID(userID uint) (uint, error) {
	var user models.User
	result := s.DB.First(&user, userID)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.RoleID, nil
}
