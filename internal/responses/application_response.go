package responses

type ApplicationAdminResponse struct {
	ApplicationID uint                      `json:"application_id"`
	User          UserResponse              `json:"user"`
	JobPost       JobPostAdminResponse      `json:"job_post"`
	Status        string                    `json:"status"`
	Answer        []UserAnswerAdminResponse `json:"answers"`
}

type ApplicationUserResponse struct {
	ApplicationID uint                `json:"application_id"`
	JobPost       JobPostUserResponse `json:"job_post"`
	Status        string              `json:"status"`
}

type UserAnswerResponse struct {
	QuestionID  uint   `json:"question_id"`
	AnswerValue string `json:"answer"`
}
type UserAnswerAdminResponse struct {
	QuestionID  uint   `json:"question_id"`
	AnswerValue string `json:"answer"`
	IsSuitable  bool   `json:"is_suitable"`
}
