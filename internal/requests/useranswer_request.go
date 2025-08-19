package requests

type UserAnswerRequest struct {
	QuestionID  uint   `json:"question_id"`
	AnswerValue string `json:"answer_value"`
}
