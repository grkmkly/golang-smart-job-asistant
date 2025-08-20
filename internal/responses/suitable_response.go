package responses

type ResponseSuitable struct {
	ApplicationID uint                      `json:"application_id"`
	User          UserResponse              `json:"user"`
	Answers       []UserAnswerAdminResponse `json:"answers"`
	IsSuitable    bool                      `json:"is_suitable"`
}
