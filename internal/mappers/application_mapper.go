package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
)

// Req Answer TO Answer Model
func ReqAnswerToAnswerModel(req requests.UserAnswerRequest) models.UserAnswer {
	return models.UserAnswer{
		AnswerValue: req.AnswerValue,
		QuestionID:  req.QuestionID,
	}
}

// Req Answers TO Answers Model
func ReqAnswersToAnswersModel(req []requests.UserAnswerRequest) []models.UserAnswer {
	var userAnswers []models.UserAnswer
	for _, uar := range req {
		answer := ReqAnswerToAnswerModel(uar)
		userAnswers = append(userAnswers, answer)
	}
	return userAnswers
}

// Req TO Model
func ReqToApplicationModel(req *requests.ApplicationRequest, userID uint) *models.Application {
	answers := ReqAnswersToAnswersModel(req.Answers)
	return &models.Application{
		Status:    "WAITING",
		JobPostID: req.JobPostID,
		UserID:    userID,
		Answer:    answers,
	}
}
