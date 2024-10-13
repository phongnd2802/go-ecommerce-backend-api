package model

type RegisterRequest struct {
	VerifyKey  string `json:"verify_key" validate:"required"`
}

type VerifyRequest struct {
	VerifyKey  string `json:"verify_key" validate:"required"`
	VerifyCode string `json:"verify_code" validate:"required"`
}

type VerifyOTPResponse struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

type SetPasswordRequest struct {
	VerifyKeyHash string `json:"verify_key_hash" validate:"required"`
	Password      string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	UserId       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}


type ForgotPasswordRequest struct {
	VerifyKey  string `json:"verify_key" validate:"required"`
}