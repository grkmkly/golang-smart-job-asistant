package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
)

// REQ TO MODEL
func ReqToJobQuestion(req *requests.JobQuestionRequest) (models.JobQuestion, error) {

	return models.JobQuestion{
		CriteriaValue:    req.CriteriaValue,
		CriteriaOperator: req.CriteriaOperator,
		QuestionOrder:    req.QuestionOrder,
		QuestionID:       req.QuestionID,
		JobPostID:        req.JobPostID,
	}, nil
}

// REQ ARRAY TO MODEL
func ReqsToJobQuestion(reqs []requests.JobQuestionRequest) ([]models.JobQuestion, error) {
	var jobQuestions []models.JobQuestion
	for _, value := range reqs {
		jobQuestion, _ := ReqToJobQuestion(&value)
		jobQuestions = append(jobQuestions, jobQuestion)
	}
	return jobQuestions, nil
}

// MODEL TO RESPONSE
func JobQuestionModelToResponse(m *models.JobQuestion) (responses.JobQuestionUserResponse, error) {
	question := QuestionToUserResponse(&m.Question)
	return responses.JobQuestionUserResponse{
		CriteriaValue:    m.CriteriaValue,
		CriteriaOperator: m.CriteriaOperator,
		QuestionID:       m.QuestionID,
		JobPostID:        m.JobPostID,
		QuestionOrder:    m.QuestionOrder,
		Question:         *question,
		//JobPost:          m.JobPost,
	}, nil
}

// MODEL ARRAY TO RESPONSE
func JobQuestionsToResponse(m []models.JobQuestion) ([]responses.JobQuestionUserResponse, error) {
	var jobQuestionsResponses []responses.JobQuestionUserResponse
	for _, jq := range m {
		jobQuestionResponse, err := JobQuestionModelToResponse(&jq)
		if err != nil {
			return nil, err
		}
		jobQuestionsResponses = append(jobQuestionsResponses, jobQuestionResponse)
	}
	return jobQuestionsResponses, nil
}
