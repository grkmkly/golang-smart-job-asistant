package requests

type JobQuestionRequest struct {
	CriteriaValue    string `json:"criteria_value"`
	CriteriaOperator string `json:"criteria_operator"`
	QuestionOrder    int    `json:"column:question_order"`
	QuestionID       uint   `json:"question_id"`

	JobPostID uint `json:"jobpost_id"`
}
