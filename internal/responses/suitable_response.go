package responses

type ResponseSuitable struct {
	ApplicationID uint               `json:"application_id"`
	User          UserResponse       `json:"user"`
	Answer        UserAnswerResponse `json:"answer"`
	IsSuitable    bool               `json:"is_suitable"`
}
