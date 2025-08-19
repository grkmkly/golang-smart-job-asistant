package responses

type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}
