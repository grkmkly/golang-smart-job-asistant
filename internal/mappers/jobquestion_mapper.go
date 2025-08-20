package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
)

// Request JobQuestion Model To JobQuestion Model
func ReqToJobQuestion(req *requests.JobQuestionRequest) (models.JobQuestion, error) {

	return models.JobQuestion{
		CriteriaValue:    req.CriteriaValue,
		CriteriaOperator: req.CriteriaOperator,
		QuestionOrder:    req.QuestionOrder,
		QuestionID:       req.QuestionID,
		JobPostID:        req.JobPostID,
	}, nil
}

// Request JobQuestion Slice To JobQuestion Model
func ReqsToJobQuestion(reqs []requests.JobQuestionRequest) ([]models.JobQuestion, error) {
	var jobQuestions []models.JobQuestion
	for _, value := range reqs {
		jobQuestion, _ := ReqToJobQuestion(&value)
		jobQuestions = append(jobQuestions, jobQuestion)
	}
	return jobQuestions, nil
}

// JobQuestion Model To Response JobQuestion Model for User
func JobQuestionModelToUserResponse(m models.JobQuestion) (responses.JobQuestionUserResponse, error) {
	question := QuestionToUserResponse(&m.Question)
	return responses.JobQuestionUserResponse{
		QuestionID:    m.QuestionID,
		QuestionOrder: m.QuestionOrder,
		Question:      *question,
	}, nil
}

// JobQuestion Slice Model To Response JobQuestion Slice Model for User
func JobQuestionsToUserResponse(m []models.JobQuestion) ([]responses.JobQuestionUserResponse, error) {
	var jobQuestionsResponses []responses.JobQuestionUserResponse

	for _, jq := range m {
		jobQuestionResponse, err := JobQuestionModelToUserResponse(jq)
		if err != nil {
			return nil, err
		}
		jobQuestionsResponses = append(jobQuestionsResponses, jobQuestionResponse)
	}
	return jobQuestionsResponses, nil
}

// JobQuestion Model To Response JobQuestion Model for Admin
func JobQuestionModelToAdminResponse(m models.JobQuestion) (responses.JobQuestionAdminResponse, error) {
	question := QuestionToAdminResponse(&m.Question)
	return responses.JobQuestionAdminResponse{
		CriteriaValue:    m.CriteriaValue,
		CriteriaOperator: m.CriteriaOperator,
		QuestionOrder:    m.QuestionOrder,
		QuestionID:       m.QuestionID,
		Question:         *question,
	}, nil
}

// JobQuestion Slice Model To Response JobQuestion Slice Model for Admin
func JobQuestionsToAdminResponse(m []models.JobQuestion) ([]responses.JobQuestionAdminResponse, error) {
	var jobQuestionsResponses []responses.JobQuestionAdminResponse

	for _, jq := range m {
		jobQuestionResponse, err := JobQuestionModelToAdminResponse(jq)
		if err != nil {
			return nil, err
		}
		jobQuestionsResponses = append(jobQuestionsResponses, jobQuestionResponse)
	}
	return jobQuestionsResponses, nil
}
