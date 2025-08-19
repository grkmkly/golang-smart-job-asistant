package responses

type JobQuestionUserResponse struct {
	QuestionOrder int `json:"question_order"`

	QuestionID uint                 `json:"question_id"`
	Question   QuestionUserResponse `json:"question"`

	//JobPostID uint `json:"job_post_id"`
}
type JobQuestionAdminResponse struct {
	CriteriaValue    string `json:"criteria_value"`
	CriteriaOperator string `json:"criteria_operator"`
	QuestionOrder    int    `json:"question_order"`

	QuestionID uint                  `json:"question_id"`
	Question   QuestionAdminResponse `json:"question"`
}
