package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/phongnd2802/go-ecommerce-backend-api/global"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/consts"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/database"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/model"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/response"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils/crypto"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils/random"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

type userAuthImpl struct{
	repo *database.Queries
}

func (ua *userAuthImpl) Register(in *model.RegisterRequest) (messageCode int, err error) {
	
	// 1. Hash Email
	fmt.Printf("Verifyket:: %s\n", in.VerifyKey)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("HashKey:: %s\n", hashKey)

	// 2. check user exists in user base
	userFound, err := ua.repo.CheckUserBaseExists(context.Background(), in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}

	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	// 3. Check OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(context.Background(), userKey).Result()
	
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed::", err)
		return response.ErrCodeInvalidOTP, err
	case otpFound != "":
		return response.ErrCodeOtpNotExists, fmt.Errorf("")
	}

	// 4. Generate OTP
	otpNew := random.GenerateSixDigit()
	fmt.Printf("OTP is::: %d\n", otpNew)

	// 5. Save OTP in Redis with expiration time
	err = global.Rdb.SetEx(context.Background(), userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTERED)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeInvalidOTP, err
	}

	// 6. Send OTP to Kafka
	body := make(map[string]interface{})
	body["otp"] = otpNew
	body["email"] = in.VerifyKey

	bodyRequest, _ := json.Marshal(body)

	message := kafka.Message{
		Key: []byte("otp-auth"),
		Value: []byte(bodyRequest),
		Time: time.Now(),
	}

	err = global.KafkaProducer.WriteMessages(context.Background(), message)
	if err != nil {
		return response.ErrCodeSendEmailOtp, err
	}

	// 7. Save OTP To MySQL
	result, err := ua.repo.InsertOTPVerify(context.Background(), database.InsertOTPVerifyParams{
		VerifyOtp: strconv.Itoa(otpNew),
		VerifyType: sql.NullInt32{Int32: 1, Valid: true},
		VerifyKey: in.VerifyKey,
		VerifyKeyHash: hashKey,
	})

	if err != nil {
		return response.ErrCodeSendEmailOtp, err
	}

	// 8. getLastId
	lastIdVerifyUser, err := result.LastInsertId()
	if err != nil {
		return response.ErrCodeSendEmailOtp, err
	}

	log.Println("LastIdVerifyUser", lastIdVerifyUser)

	return response.CodeSuccess, nil
}

func (ua *userAuthImpl) Login() (messageCode int, err error) {
	return response.CodeSuccess, nil
}

func (ua *userAuthImpl) VerifyOTP(in *model.VerifyRequest) (messageCode int, out model.VerifyOTPResponse, err error) {
	hashKey := crypto.GetHash(in.VerifyKey)

	// get otp
	otpFound, err := global.Rdb.Get(context.Background(), utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return response.ErrCodeInvalidOTP, out, err
	}

	if in.VerifyCode != otpFound {
		return response.ErrCodeInvalidOTP, out, fmt.Errorf("OTP not match")
	}

	infoOtp, err := ua.repo.GetInfoOTP(context.Background(), hashKey)
	if err != nil {
		return response.ErrCodeInvalidOTP, out, err
	}

	// Update status verified
	err = ua.repo.UpdateUserVerificationStatus(context.Background(), hashKey)
	if err != nil {
		return response.ErrCodeUpdateVerifiedStatus, out, err
	}

	// Response
	out.Token = infoOtp.VerifyKeyHash
	out.UserId = strconv.Itoa(int(infoOtp.VerifyID))
	
	return response.CodeSuccess, out, nil
}

func (ua *userAuthImpl) UpdatePasswordVerified() (messageCode int,err error) {
	return response.CodeSuccess, nil
}


func NewUserAuthImpl(repo *database.Queries) *userAuthImpl {
	return &userAuthImpl{
		repo: repo,
	}
}



