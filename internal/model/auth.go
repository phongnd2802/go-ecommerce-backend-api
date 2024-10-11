package model

type RegisterRequest struct {
	VerifyKey  string `json:"verify_key" validate:"required"`
	VerifyType string `json:"verify_type" validate:"required"`
}

type VerifyRequest struct {
	VerifyKey  string `json:"verify_key" validate:"required"`
	VerifyCode string `json:"verify_code" validate:"required"`
}

type VerifyOTPResponse struct {
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}
