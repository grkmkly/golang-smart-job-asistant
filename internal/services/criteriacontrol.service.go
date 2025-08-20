package services

import (
	"smartjob/internal/mappers"
	"smartjob/internal/models"
	"smartjob/internal/responses"
	"strings"

	"gorm.io/gorm"
)

type CriteriaControlService struct {
	DB *gorm.DB
}

func NewCriteriaControlService(db *gorm.DB) *CriteriaControlService {
	return &CriteriaControlService{
		DB: db,
	}
}

func (s *CriteriaControlService) CriteriaControl(application models.Application) []responses.ResponseSuitable {
	var answers []responses.UserAnswerAdminResponse

	for _, a := range application.Answer {
		answer := mappers.UserAnswerAdminToResponse(&a)
		answer.IsSuitable = s.checkAnswerSuitability(&a, application.JobPost.JobQuestions)
		answers = append(answers, *answer)
		continue

	}

	appIsSuitable := true
	for _, answer := range answers {
		if !answer.IsSuitable {
			appIsSuitable = false
			break
		}
	}

	response := responses.ResponseSuitable{
		ApplicationID: application.ID,
		User:          *mappers.UserModelToResponse(&application.User),
		Answers:       answers,
		IsSuitable:    appIsSuitable,
	}
	return []responses.ResponseSuitable{response}
}

func (s *CriteriaControlService) checkAnswerSuitability(userAnswer *models.UserAnswer, jobQuestions []models.JobQuestion) bool {
	for _, jq := range jobQuestions {
		if jq.QuestionID == userAnswer.QuestionID {
			if strings.Contains(strings.ToLower(jq.Question.Type), "dropdown") {
				return strings.EqualFold(jq.CriteriaValue, userAnswer.AnswerValue)
			}
			return true
		}
	}
	return true
}
