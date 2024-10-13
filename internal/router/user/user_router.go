package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/controller/account"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/middleware"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/response"
)

type UserRouter struct {}

func (ur *UserRouter) initUserRouter(Router *mux.Router) {
	userPublicRouter := Router.PathPrefix("/user").Subrouter()
	// Router Public
	userPublicRouter.HandleFunc("/register", account.Auth.Register).Methods("POST") 
	userPublicRouter.HandleFunc("/otp", account.Auth.VerifyOTP).Methods("POST")
	userPublicRouter.HandleFunc("/set_password", account.Auth.UpdatePasswordRegister).Methods("POST")
	userPublicRouter.HandleFunc("/login", account.Auth.Login).Methods("POST")
	userPublicRouter.HandleFunc("/forgot_password", account.Auth.ForgotPassword).Methods("POST")

	////////////////////////////////////////////////////////////////////////////////
	userPrivateRouter := Router.PathPrefix("/user").Subrouter()

	// Middleware
	userPrivateRouter.Use(middleware.Authentication) // Authentication
	userPrivateRouter.Use(middleware.Permission) // Permisison

	// Router Private
	userPrivateRouter.HandleFunc("/get_info", func(w http.ResponseWriter, r *http.Request) {
		response.SuccessResponse(w, response.CodeSuccess, "Private User")
	}).Methods("GET")
	userPrivateRouter.HandleFunc("/logout", account.Auth.Logout).Methods("POST")
}