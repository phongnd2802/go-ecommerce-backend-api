package initialize

import (
	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/database"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/service"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/service/impl"
)


func InitServiceInterface() {
	queries := database.New(global.Mdb)
	// User Service Interface
	service.InitUserAuth(impl.NewUserAuthImpl(queries))
}