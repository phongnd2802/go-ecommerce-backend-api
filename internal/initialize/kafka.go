package initialize

import (
	"fmt"
	"log"
	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/segmentio/kafka-go"
)

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
	fmt.Println("Connected to Kafka")
	global.Logger.Info("Init Kafka Success")
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("failed to close kafka producer: %v", err)
	}
}
