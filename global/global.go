package global

import (
	"database/sql"

	"github.com/minio/minio-go/v7"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/logger"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *sql.DB
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
	MinioClient   *minio.Client
)
