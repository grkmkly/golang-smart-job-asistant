package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/responses"
)

// User Answer Model To Response Answer for User
func UserAnswerToResponse(answer *models.UserAnswer) *responses.UserAnswerResponse {
	if answer == nil {
		return nil
	}
	response := &responses.UserAnswerResponse{
		AnswerValue: answer.AnswerValue,
		QuestionID:  answer.QuestionID,
	}
	return response
}

// UserAnswer Model To Response Answer Slice for User
func UserAnswersModelToResponseSlice(answers []models.UserAnswer) ([]responses.UserAnswerResponse, error) {
	var userAnswerResponses []responses.UserAnswerResponse
	for _, answer := range answers {
		response := UserAnswerToResponse(&answer)
		userAnswerResponses = append(userAnswerResponses, *response)
	}
	return userAnswerResponses, nil
}

// UserAnswer Model To Response Answer for Admin
func UserAnswerAdminToResponse(answer *models.UserAnswer) *responses.UserAnswerAdminResponse {
	if answer == nil {
		return nil
	}
	response := &responses.UserAnswerAdminResponse{
		AnswerValue: answer.AnswerValue,
		QuestionID:  answer.QuestionID,
		IsSuitable:  true,
	}
	return response
}

// UserAnswer Model To Response Answer Slice for Admin
func UserAnswersAdminToResponseSlice(answers []models.UserAnswer) ([]responses.UserAnswerAdminResponse, error) {
	var userAnswerAdminResponses []responses.UserAnswerAdminResponse
	for _, answer := range answers {
		response := UserAnswerAdminToResponse(&answer)
		userAnswerAdminResponses = append(userAnswerAdminResponses, *response)
	}
	return userAnswerAdminResponses, nil
}
