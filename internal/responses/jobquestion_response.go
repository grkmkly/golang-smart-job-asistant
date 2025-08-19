package responses

type JobQuestionUserResponse struct {
	CriteriaValue    string `gorm:"type:varchar(255);column:criteria_value"`
	CriteriaOperator string `gorm:"type:varchar(50);column:criteria_operator"`
	QuestionOrder    int    `gorm:"column:question_order"`

	QuestionID uint                 `gorm:"not null"`
	Question   QuestionUserResponse `gorm:"foreignKey:QuestionID"`

	JobPostID uint `gorm:"not null"`
	//	JobPost   models.JobPost `gorm:"foreignKey:JobPostID"`
}
