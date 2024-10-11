package response

const (
	CodeSuccess = 2000

	ErrCodeParamInvalid         = 6000
	ErrCodeUserHasExists        = 6001
	ErrCodeInvalidOTP           = 6002
	ErrCodeOtpNotExists         = 6003
	ErrCodeSendEmailOtp         = 6004
	ErrCodeUpdateVerifiedStatus = 6005
)

var msg = map[int]string{
	CodeSuccess:                 "Success",
	ErrCodeUserHasExists:        "Email already exist",
	ErrCodeInvalidOTP:           "OTP invalid",
	ErrCodeOtpNotExists:         "OTP exists but not registered",
	ErrCodeSendEmailOtp:         "Send Email OTP error",
	ErrCodeParamInvalid:         "Param is invalid",
	ErrCodeUpdateVerifiedStatus: "Err Verified Update",
}
