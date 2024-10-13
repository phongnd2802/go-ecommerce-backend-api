package initialize

import (
	// "fmt"
	// "log"
	// "net/http"

	"github.com/gorilla/mux"
	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)


func Run() *mux.Router {
	// Config
	LoadConfig()

	// Init Logger
	InitLogger()
	global.Logger.Info("Config Log OK!", zap.String("ok", "success"))

	// Init Mysql databasae
	InitMysql()

	// Init Seed Data
	InitSeedData()

	// Init Service Interface
	InitServiceInterface()

	// Init Redis Cache
	InitRedis()

	// Init Kafka
	InitKafka()

	// Init MinIO
	InitMinIO()

	// Init Router
	r := InitRouter()

	return r
}