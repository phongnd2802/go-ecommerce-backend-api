package initialize

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/phongnd2802/go-ecommerce-backend-api/global"
)


func InitMinIO() {
	accessKeyId := "yuNxswxfaBY2bGigPyzR"
	secretAccessKey := "1kianO1pGfgtPcijg38KPuqYiTwZ6z2s1QNUkkj7"
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: true,
	})

	if err != nil {
		global.Logger.Error("Connect MinIO failed")
		log.Fatalln(err)
	}
	global.MinioClient = minioClient
	fmt.Println("Connected To MinIO")
	global.Logger.Info("Init MinIO Success")
}