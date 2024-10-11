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

	// Init Service Interface
	InitServiceInterface()

	// Init Redis Cache
	InitRedis()

	// Init Kafka
	InitKafka()


	// Init Router
	r := InitRouter()

	// log.Println("Server started on port 8000")
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", global.Config.Server.Port), r))

	return r
}