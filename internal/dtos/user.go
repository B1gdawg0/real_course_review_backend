package dtos

type UserShortResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type AuthRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name"`
}

type AuthResponse struct {
	Token string            `json:"token"`
	User  UserShortResponse `json:"user"`
}