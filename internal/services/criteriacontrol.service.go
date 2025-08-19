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
	var responsesA []responses.ResponseSuitable
	var response responses.ResponseSuitable

	for _, jq := range application.JobPost.JobQuestions {
		var answer *models.UserAnswer
		for _, a := range application.Answer {
			if a.QuestionID == jq.QuestionID {
				answer = &a
				break
			}
		}
		if strings.Contains(strings.ToLower(jq.Question.Type), "dropdown") {
			if strings.EqualFold(jq.CriteriaValue, answer.AnswerValue) {
				response = responses.ResponseSuitable{
					ApplicationID: application.ID,
					User:          *mappers.UserModelToResponse(&application.User),
					Answer:        *mappers.UserAnswerToResponse(answer),
					IsSuitable:    true,
				}
			} else {
				response = responses.ResponseSuitable{
					ApplicationID: application.ID,
					User:          *mappers.UserModelToResponse(&application.User),
					Answer:        *mappers.UserAnswerToResponse(answer),
					IsSuitable:    false,
				}
			}
			responsesA = append(responsesA, response)
			continue
		} else {
			response = responses.ResponseSuitable{
				ApplicationID: application.ID,
				User:          *mappers.UserModelToResponse(&application.User),
				Answer:        *mappers.UserAnswerToResponse(answer),
				IsSuitable:    true,
			}
			responsesA = append(responsesA, response)
			continue
		}
	}
	return responsesA
}
