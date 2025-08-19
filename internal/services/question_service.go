package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"

	"gorm.io/gorm"
)

type QuestionService struct {
	Db *gorm.DB
}

func NewQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{db}
}

func (s *QuestionService) CreateWithOption(req *requests.QuestionRequest, userID uint) error {
	question := mappers.QuestionRequestToQuestion(req, userID)

	if err := s.Db.Create(question).Error; err != nil {
		return err
	}
	return nil
}

func (s *QuestionService) GetUserWithOption() ([]responses.QuestionUserResponse, error) {
	var questions []models.Question
	err := s.Db.Preload("Options").
		Preload("CreatedBy").
		Order("created_at DESC").
		Find(&questions).Error
	if err != nil {
		return nil, err
	}
	responseQuestion := mappers.QuestionsToUserResponse(questions)
	return responseQuestion, nil

}
func (s *QuestionService) GetAdminWithOption() ([]responses.QuestionAdminResponse, error) {
	var questions []models.Question
	err := s.Db.Preload("Options").
		Preload("CreatedBy").
		Order("created_at DESC").
		Find(&questions).Error
	if err != nil {
		return nil, err
	}
	responseQuestion := mappers.QuestionsToAdminResponse(questions)
	return responseQuestion, nil
}
