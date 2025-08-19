package requests

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	FirstName   string `json:"first_name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number"  binding:"required"`
}
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
