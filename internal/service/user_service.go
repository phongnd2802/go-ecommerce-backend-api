package service

import "github.com/phongnd2802/go-ecommerce-backend-api/internal/model"

type (
	IUserAuth interface {
		Register(in *model.RegisterRequest) (messageCode int, err error)
		Login() (messageCode int, err error)
		VerifyOTP(in *model.VerifyRequest) (messageCode int, out model.VerifyOTPResponse, err error)
		UpdatePasswordVerified() (messageCode int, err error)
	}

	IUserInfo interface {
		GetInfo() (messageCode int, err error)
	}	

	IUserAdmin interface {
		FindUserByAdmin() (messageCode int, err error)
	}
)

var (
	iUserAuth IUserAuth
	iUserInfo IUserInfo
	iUserAdmin IUserAdmin
)

func UserAuth() IUserAuth {
	if iUserAuth == nil {
		panic("implement iUserAuth not found for interface IUserAuth")
	}
	return iUserAuth
}


func InitUserAuth(i IUserAuth) {
	iUserAuth = i
}

func UserInfo() IUserInfo {
	if iUserInfo == nil {
		panic("implement iUserInfo not found for interface IUserInfo")
	}
	return iUserInfo
}

func InitUserInfo(i IUserInfo) {
	iUserInfo = i
}

func UserAdmin() IUserAdmin {
	if iUserAdmin == nil {
		panic("implement iUserAuth not found for interface IUserAdmin")
	}
	return iUserAdmin
}
