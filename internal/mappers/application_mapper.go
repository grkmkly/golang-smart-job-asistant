package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
)

// Request Answer To Answer Model
func ReqAnswerToAnswerModel(req requests.UserAnswerRequest) models.UserAnswer {
	return models.UserAnswer{
		AnswerValue: req.AnswerValue,
		QuestionID:  req.QuestionID,
	}
}

// Request Answer Slice To Answers Model Slice
func ReqAnswersToAnswersModel(req []requests.UserAnswerRequest) []models.UserAnswer {
	var userAnswers []models.UserAnswer
	for _, uar := range req {
		answer := ReqAnswerToAnswerModel(uar)
		userAnswers = append(userAnswers, answer)
	}
	return userAnswers
}

// Request Application To Model Application
func ReqToApplicationModel(req *requests.ApplicationRequest, userID uint, postID uint) *models.Application {
	answers := ReqAnswersToAnswersModel(req.Answers)
	return &models.Application{
		Status:    "WAITING",
		JobPostID: postID,
		UserID:    userID,
		Answer:    answers,
	}
}

// Application Model To Response Application for Admin
func ApplicationModelToResponse(application *models.Application) (*responses.ApplicationAdminResponse, error) {
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
	response := &responses.ApplicationAdminResponse{
		ApplicationID: application.ID,
		JobPost:       *jobPost,
		User:          *UserModelToResponse(&application.User),
		Answer:        answer,
		Status:        application.Status,
	}
	return response, nil
}

// Application Slice Model To Response Application Slice for Admin
func ApplicationModelsToAdminResponseSlice(applications []models.Application) ([]responses.ApplicationAdminResponse, error) {
	var responses []responses.ApplicationAdminResponse
	for _, application := range applications {
		response, err := ApplicationModelToResponse(&application)
		if err != nil {
			return nil, err
		}
		responses = append(responses, *response)
	}
	return responses, nil
}

// Application Model To Response Application for User
func ApplicationModelToUserResponse(application *models.Application) *responses.ApplicationUserResponse {
	return &responses.ApplicationUserResponse{
		ApplicationID: application.ID,
		Status:        application.Status,
		JobPost:       *JobPostModelToUserResponse(&application.JobPost),
	}
}

// Application Slice Model To Response Application Slice for User
func ApplicationModelsToUserResponseSlice(applications []models.Application) []responses.ApplicationUserResponse {
	var userResponses []responses.ApplicationUserResponse
	for _, app := range applications {
		response := ApplicationModelToUserResponse(&app)
		userResponses = append(userResponses, *response)
	}
	return userResponses
}
