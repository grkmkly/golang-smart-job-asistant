package responses

type ResponseSuitable struct {
	ApplicationID uint                      `json:"application_id"`
	User          UserResponse              `json:"user"`
	JobPost       JobPostAdminResponse      `json:"job_post"`
	Status        string                    `json:"status"`
	Answers       []UserAnswerAdminResponse `json:"answers"`
	IsSuitable    bool                      `json:"is_suitable"`
}
