package mappers

import (
	"smartjob/internal/models"
	"smartjob/internal/requests"
	"smartjob/internal/responses"
)

// REQ TO MODEL
func JobPostRequestToJobPost(req *requests.JobPostRequest, userID uint) (*models.JobPost, error) {
	jobQuestions, err := ReqsToJobQuestion(req.JobQuestion)
	if err != nil {
		return nil, err
	}

	return &models.JobPost{
		Title:        req.Title,
		Content:      req.Content,
		EndAt:        req.EndAt,
		IsActive:     true,
		CreatedByID:  userID,
		JobQuestions: jobQuestions,
	}, err
}

// MODEL TO RESPONSE
func JobPostModelToUserResponse(jobPost *models.JobPost) (*responses.JobPostUserResponse, error) {
	//jobQuestions, err := JobQuestionsToResponse(jobPost.JobQuestions)
	//if err != nil {
	//	return nil, err
	//}
	return &responses.JobPostUserResponse{
		Title:       jobPost.Title,
		Content:     jobPost.Content,
		EndAt:       jobPost.EndAt,
		IsActive:    jobPost.IsActive,
		CreatedByID: jobPost.CreatedByID,
		//JobQuestions: jobQuestions,
	}, nil
}

func JobPostModelToUserResponseSlice(jobPosts *[]models.JobPost) ([]responses.JobPostUserResponse, error) {
	var responses []responses.JobPostUserResponse
	for _, jobPost := range *jobPosts {
		response, err := JobPostModelToUserResponse(&jobPost)
		if err != nil {
			return nil, err
		}
		responses = append(responses, *response)
	}
	return responses, nil
}

func JobPostModelToAdminResponse(jobPost *models.JobPost) (*responses.JobPostAdminResponse, error) {
	jobQuestions, err := JobQuestionsToAdminResponse(jobPost.JobQuestions)
	if err != nil {
		return nil, err
	}
	return &responses.JobPostAdminResponse{
		Title:        jobPost.Title,
		Content:      jobPost.Content,
		EndAt:        jobPost.EndAt,
		IsActive:     jobPost.IsActive,
		CreatedByID:  jobPost.CreatedByID,
		JobQuestions: jobQuestions,
	}, nil
}

func JobPostModelToAdminResponseSlice(jobPosts *[]models.JobPost) ([]responses.JobPostAdminResponse, error) {
	var responses []responses.JobPostAdminResponse
	for _, jobPost := range *jobPosts {
		response, err := JobPostModelToAdminResponse(&jobPost)
		if err != nil {
			return nil, err
		}
		responses = append(responses, *response)
	}
	return responses, nil
}
