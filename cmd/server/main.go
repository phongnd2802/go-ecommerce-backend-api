package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/initialize"
	_ "github.com/phongnd2802/go-ecommerce-backend-api/cmd/swag/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Go Ecommerce Backend API
// @version 1.0
// @description This is documentation for API
// @termsOfService github.com/phongnd2802/go-ecommerce-backend-api

// @contact.name PhongND
// @contact.url https://github.com/phongnd2802/go-ecommerce-backend-api
// @contact.email duyphong02802@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
func main() {
	r := initialize.Run()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started on port 8000")
	http.ListenAndServe(fmt.Sprintf(":%d", global.Config.Server.Port), r)
}