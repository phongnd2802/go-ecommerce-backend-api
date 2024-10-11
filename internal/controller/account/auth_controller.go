package account

import (
	"net/http"

	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/model"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/service"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/response"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils"
	"go.uber.org/zap"
)

type cUserAuth struct{}

var Auth = new(cUserAuth)

// User Registration documentation
// @Summary      User Registration
// @Description  When user is registered send otp to email
// @Tags         Account Management
// @Accept       json
// @Produce      json
// @Param		 payload body model.RegisterRequest true "payload"
// @Success      200	{object}	response.ResponseData
// @Failure		 400	{object}	response.ResponseData
// @Failure 	 500	{object}	response.ResponseData
// @Router       /user/register [post]
func (c *cUserAuth) Register(w http.ResponseWriter, r *http.Request) {
	var params model.RegisterRequest
	if err := utils.ParseJSON(r, &params); err != nil {
		response.ErrorResponse(w, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeStatus, err := service.UserAuth().Register(&params)
	if err != nil {
		global.Logger.Error("Error registering user", zap.Error(err))
		response.ErrorResponse(w, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(w, codeStatus, nil)
}



// User Verify OTP documentation
// @Summary      User Verify OTP
// @Description  When User Verify OTP then Update Status 
// @Tags         Account Management
// @Accept       json
// @Produce      json
// @Param		 payload body model.VerifyRequest true "payload"
// @Success      200	{object}	response.ResponseData
// @Failure		 400	{object}	response.ResponseData
// @Failure 	 500	{object}	response.ResponseData
// @Router       /user/otp [post]
func (c *cUserAuth) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var params model.VerifyRequest
	if err := utils.ParseJSON(r, &params); err != nil {
		response.ErrorResponse(w, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeStatus, data, err := service.UserAuth().VerifyOTP(&params)
	if err != nil {
		global.Logger.Error("Error Verify OTP", zap.Error(err))
		response.ErrorResponse(w, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(w, codeStatus, data)
}



// User Set Passaword Verified documentation
// @Summary      User Set Password Verified
// @Description  Set Password when user verified
// @Tags         Account Management
// @Accept       json
// @Produce      json
// @Param		 payload body model.SetPasswordRequest true "payload"
// @Success      200	{object}	response.ResponseData
// @Failure		 400	{object}	response.ResponseData
// @Failure 	 500	{object}	response.ResponseData
// @Router       /user/set_password [post]
func (c *cUserAuth) UpdatePasswordVerified(w http.ResponseWriter, r *http.Request) {
	var params model.SetPasswordRequest
	if err := utils.ParseJSON(r, &params); err != nil {
		response.ErrorResponse(w, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeStatus, err := service.UserAuth().UpdatePasswordVerified(&params)
	if err != nil {
		global.Logger.Error("Error Set Password", zap.Error(err))
		response.ErrorResponse(w, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(w, codeStatus, nil)
}

// User Login documentation
// @Summary      User Login
// @Description  User Login
// @Tags         Account Management
// @Accept       json
// @Produce      json
// @Param		 payload body model.LoginRequest true "payload"
// @Success      200	{object}	response.ResponseData
// @Failure		 400	{object}	response.ResponseData
// @Failure 	 500	{object}	response.ResponseData
// @Router       /user/login [post]
func (c *cUserAuth) Login(w http.ResponseWriter, r *http.Request) {
	var params model.LoginRequest
	if err := utils.ParseJSON(r, &params); err != nil {
		response.ErrorResponse(w, response.ErrCodeParamInvalid, err.Error())
		return
	}
	codeStatus, data, err := service.UserAuth().Login(&params)
	if err != nil {
		global.Logger.Error("Error Login", zap.Error(err))
		response.ErrorResponse(w, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(w, codeStatus, data)
}