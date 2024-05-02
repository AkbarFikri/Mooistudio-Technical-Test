package dto

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required,max=255"`
}

type AuthResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}

type UserTokenData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
