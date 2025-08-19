package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
	"smartjob/internal/utils"
)

// REQ TO MODEL
func RegisterRequestToUser(req requests.RegisterRequest) (models.User, error) {

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.User{}, err
	}
	newUser := models.User{
		FirstName:    req.FirstName,
		Surname:      req.Surname,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		PhoneNumber:  req.PhoneNumber,
		RoleID:       1,
		IsActive:     true,
	}
	return newUser, nil
}

func UserModelToResponse(user *models.User) *responses.UserResponse {
	if user == nil {
		return nil
	}

	response := &responses.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		Surname:   user.Surname,
		Email:     user.Email,
	}
	return response
}
