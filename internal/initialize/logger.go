package initialize

import (
	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/logger"
)


func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}