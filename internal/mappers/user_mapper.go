package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
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
