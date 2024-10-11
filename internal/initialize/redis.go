package initialize

import (
	"context"
	"fmt"

	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)


func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: "",
		DB:       0,
		Protocol: r.PoolSize,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("Redis init error", zap.Error(err))
	}
	fmt.Println("Redis is running")
	global.Logger.Info("Redis init success")
	global.Rdb = rdb
}