package dto

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type AuthResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	User    User   `json:"user"`
	Token   string `json:"token"`
}
