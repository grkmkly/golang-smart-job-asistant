package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
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
func ReqToApplicationModel(req *requests.ApplicationRequest, userID uint, postID uint) *models.Application {
	answers := ReqAnswersToAnswersModel(req.Answers)
	return &models.Application{
		Status:    "WAITING",
		JobPostID: postID,
		UserID:    userID,
		Answer:    answers,
	}
}

// APPLICATION MODEL TO RESPONSE
func ApplicationModelToResponse(application *models.Application) (*responses.ApplicationResponse, error) {
	if application == nil {
		return nil, nil
	}
	jobPost, err := JobPostModelToAdminResponse(&application.JobPost)
	if err != nil {
		return nil, err
	}
	answer, err := UserAnswersAdminToResponseSlice(application.Answer)
	if err != nil {
		return nil, err
	}
	response := &responses.ApplicationResponse{
		ApplicationID: application.ID,
		JobPost:       *jobPost,
		User:          *UserModelToResponse(&application.User),
		Answer:        answer,
	}
	return response, nil
}

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

// USER ANSWERS MODELS TO RESPONSES
func UserAnswersModelToResponseSlice(answers []models.UserAnswer) ([]responses.UserAnswerResponse, error) {
	var userAnswerResponses []responses.UserAnswerResponse
	for _, answer := range answers {
		response := UserAnswerToResponse(&answer)
		userAnswerResponses = append(userAnswerResponses, *response)
	}
	return userAnswerResponses, nil
}

func UserAnswerAdminToResponse(answer *models.UserAnswer) *responses.UserAnswerAdminResponse {
	if answer == nil {
		return nil
	}
	response := &responses.UserAnswerAdminResponse{
		AnswerValue: answer.AnswerValue,
		QuestionID:  answer.QuestionID,
	}
	return response
}

func UserAnswersAdminToResponseSlice(answers []models.UserAnswer) ([]responses.UserAnswerAdminResponse, error) {
	var userAnswerAdminResponses []responses.UserAnswerAdminResponse
	for _, answer := range answers {
		response := UserAnswerAdminToResponse(&answer)
		userAnswerAdminResponses = append(userAnswerAdminResponses, *response)
	}
	return userAnswerAdminResponses, nil
}

// APPLICATION MODELS TO RESPONES
func ApplicationModelsToResponseSlice(applications []models.Application) ([]responses.ApplicationResponse, error) {
	var responses []responses.ApplicationResponse
	for _, application := range applications {
		response, err := ApplicationModelToResponse(&application)
		if err != nil {
			return nil, err
		}
		responses = append(responses, *response)
	}
	return responses, nil
}
