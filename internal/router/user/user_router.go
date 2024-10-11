package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/controller/account"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/response"
)

type UserRouter struct {}

func (ur *UserRouter) InitUserRouter(Router *mux.Router) {
	userPublicRouter := Router.PathPrefix("/user").Subrouter()
	// userPublicRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	response.WriteJSON(w, http.StatusOK, map[string]interface{}{
	// 		"message": "user",
	// 	})
	// })
	userPublicRouter.HandleFunc("/register", account.Auth.Register).Methods("POST") 
	userPublicRouter.HandleFunc("/otp", account.Auth.VerifyOTP).Methods("POST")
	
	userPrivateRouter := Router.PathPrefix("/user").Subrouter()
	//userPrivateRouter.Use()
	userPrivateRouter.HandleFunc("/get_info", func(w http.ResponseWriter, r *http.Request) {
		response.SuccessResponse(w, response.CodeSuccess, "Private User")
	}).Methods("GET")
}