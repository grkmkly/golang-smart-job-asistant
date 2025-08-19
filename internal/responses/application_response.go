package responses

type ApplicationResponse struct {
	ApplicationID uint                      `json:"application"`
	User          UserResponse              `json:"user"`
	JobPost       JobPostAdminResponse      `json:"job_post"`
	Answer        []UserAnswerAdminResponse `json:"answers"`
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
