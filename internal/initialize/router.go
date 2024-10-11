package initialize

import (
	"github.com/MadAppGang/httplog"
	"github.com/gorilla/mux"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/controller"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/router"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(httplog.Logger)
	userRouter := router.RouterApp.User

	MainRouter := r.PathPrefix("/api/v1").Subrouter()
	MainRouter.HandleFunc("/", controller.Pong).Methods("GET")
	userRouter.InitUserRouter(MainRouter)
	return r
}