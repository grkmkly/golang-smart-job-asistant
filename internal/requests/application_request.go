package requests

type ApplicationRequest struct {
	JobPostID uint                `json:"job_post_id"`
	Answers   []UserAnswerRequest `json:"answers"`
}
