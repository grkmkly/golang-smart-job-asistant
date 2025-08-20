package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
)

// Question Request To Question Model
func QuestionRequestToQuestion(req *requests.QuestionRequest, userID uint) *models.Question {
	var options []models.QuestionOption

	if req.Options != nil {
		for i, optReq := range req.Options {
			options = append(options, models.QuestionOption{
				OptionValue: optReq.OptionValue,
				Priority:    len(req.Options) - i,
			})
		}
	}
	return &models.Question{
		Content:     req.Content,
		CreatedByID: userID,
		Type:        req.Type,
		Options:     options,
	}
}

// Question Model To Response Question for User
func QuestionToUserResponse(q *models.Question) *responses.QuestionUserResponse {

	var options []responses.OptionResponse
	for _, opt := range q.Options {
		option := responses.OptionResponse{
			ID:    int(opt.ID),
			Value: opt.OptionValue,
		}

		options = append(options, option)
	}
	return &responses.QuestionUserResponse{
		Content: q.Content,
		Type:    q.Type,
		Options: options,
	}
}

// Question Slice To Response Question Slice for User
func QuestionsToUserResponse(q []models.Question) []responses.QuestionUserResponse {
	var questions []responses.QuestionUserResponse

	for _, q := range q {
		question := QuestionToUserResponse(&q)
		questions = append(questions, *question)
	}
	return questions
}

// Question Model To Response Question for Admin
func QuestionToAdminResponse(q *models.Question) *responses.QuestionAdminResponse {
	var options []responses.OptionResponse
	for _, opt := range q.Options {
		option := responses.OptionResponse{
			ID:    int(opt.ID),
			Value: opt.OptionValue,
		}

		options = append(options, option)
	}
	return &responses.QuestionAdminResponse{
		Content:   q.Content,
		Type:      q.Type,
		Options:   options,
		CreatedBy: *UserModelToResponse(&q.CreatedBy),
	}
}

// Question Slice Model To Response Question Slice for Admin
func QuestionsToAdminResponse(q []models.Question) []responses.QuestionAdminResponse {
	var questions []responses.QuestionAdminResponse

	for _, q := range q {
		question := QuestionToAdminResponse(&q)
		questions = append(questions, *question)
	}
	return questions
}
