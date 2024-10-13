package initialize

import (
	"github.com/MadAppGang/httplog"
	"github.com/gorilla/mux"
	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/controller"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/router"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	if global.Config.Server.Mode == "dev" {
		r.Use(httplog.Logger)
	}

	MainRouter := r.PathPrefix("/api/v1").Subrouter()
	// Monitor Check
	MainRouter.HandleFunc("/", controller.Pong).Methods("GET")

	userRouter := router.RouterApp.User
	ownerRouter := router.RouterApp.Owner
	adminRouter := router.RouterApp.Admin

	// Init Router for User
	userRouter.InitRouter(MainRouter)

	// Init Router for Owner
	ownerRouter.InitRouter(MainRouter)

	// Init Router for Admin
	adminRouter.InitRouter(MainRouter)
	return r
}