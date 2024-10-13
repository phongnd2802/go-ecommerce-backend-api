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
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/jwt"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/response"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils/crypto"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils/random"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
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
	userKey := utils.GetUserKeyOTP(hashKey)
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


func (ua *userAuthImpl) Login(in *model.LoginRequest) (int, *model.LoginResponse, error) {
	// Check Email
	userFound, err := ua.repo.GetUserBase(context.Background(), in.Email)
	if err != nil {
		return response.ErrCodeEmailNotVerifiedOrNotRegistered, nil, err
	}

	// Check Password
	matched := crypto.CheckPasswordWithSalt(in.Password, userFound.UserSalt, userFound.UserPassword)
	if !matched {
		return response.ErrCodePasswordDoNotMatch, nil, err
	}

	// Generate AccessToken, RefreshToken
	publicKeyStr, privateKey, err := crypto.GenerateRSAKeyPair(1024)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	payload := map[string]interface{}{
		"email": userFound.UserAccount,
		"uid": userFound.UserID,
	}
	accessToken, err := jwt.GenerateToken(payload, privateKey, int64(consts.TIME_ACCESSS_TOKEN))
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	refreshToken, err := jwt.GenerateToken(payload, privateKey, consts.TIME_REFRESH_TOKEN)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	// Save accessToken and PublicKey in Redis
	hashEmail := crypto.GetHash(userFound.UserAccount)
	userKeyToken := utils.GetUserKeyToken(hashEmail)
	err = global.Rdb.SetEx(context.Background(), userKeyToken, accessToken, time.Duration(consts.TIME_ACCESSS_TOKEN) * time.Second).Err()
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	userKeySecret := utils.GetUserKeySecret(hashEmail)
	err = global.Rdb.SetEx(context.Background(), userKeySecret, publicKeyStr, time.Duration(consts.TIME_ACCESSS_TOKEN) * time.Second).Err()
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}


	// Save refreshToken and publicKey in MySQL
	fmt.Println("PublicKey:>>> ", publicKeyStr)
	_, err = ua.repo.InsertUserToken(context.Background(), database.InsertUserTokenParams{
		RefreshToken: refreshToken,
		PublicKey: publicKeyStr,
		UserID: userFound.UserID,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	// Update state login MySQL
	err = ua.repo.UpdateInfoLogin(context.Background(), database.UpdateInfoLoginParams{
		UserID: userFound.UserID,
		UserLoginIp: sql.NullString{String: "0.0.0.0", Valid: true},
	})
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}


	// Response
	out := &model.LoginResponse{
		UserId: int(userFound.UserID),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
	return response.CodeSuccess, out, nil
}


func (ua *userAuthImpl) VerifyOTP(in *model.VerifyRequest) (int, *model.VerifyOTPResponse, error) {
	hashKey := crypto.GetHash(in.VerifyKey)

	// get otp
	otpFound, err := global.Rdb.Get(context.Background(), utils.GetUserKeyOTP(hashKey)).Result()
	if err != nil {
		return response.ErrCodeInvalidOTP, nil, err
	}

	if in.VerifyCode != otpFound {
		return response.ErrCodeInvalidOTP, nil, fmt.Errorf("OTP not match")
	}

	infoOtp, err := ua.repo.GetInfoOTP(context.Background(), hashKey)
	if err != nil {
		return response.ErrCodeInvalidOTP, nil, err
	}

	// Update status verified
	err = ua.repo.UpdateUserVerificationStatus(context.Background(), hashKey)
	if err != nil {
		return response.ErrCodeUpdateVerifiedStatus, nil, err
	}

	// Response
	out := &model.VerifyOTPResponse{
		Token: infoOtp.VerifyKeyHash,
		UserId: int(infoOtp.VerifyID),
	}
	
	return response.CodeSuccess, out, nil
}

func (ua *userAuthImpl) UpdatePasswordRegister(in *model.SetPasswordRequest) (messageCode int, err error) {
	// Check Verified
	infoVerify, err := ua.repo.GetValidVerified(context.Background(), in.VerifyKeyHash)
	if err != nil {
		return response.ErrCodeEmailNotVerified, err
	}

	// Random Salt
	ranSalt, err := random.RandomSalt(16)
	if err != nil {
		return response.ErrCodeInternalServerError, err
	}

	// Hash Password
	hashedPassword, err := crypto.HashPasswordWithSalt(in.Password, ranSalt)
	if err != nil {
		return response.ErrCodeInternalServerError, err
	}

	// Insert Info Account To User Base MySQL
	resultBase, err := ua.repo.InsertUserBase(context.Background(), database.InsertUserBaseParams{
		UserAccount: infoVerify.VerifyKey,
		UserPassword: hashedPassword,
		UserSalt: ranSalt,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, err
	}
	lastIdUserBase, err := resultBase.LastInsertId()
	if err != nil {
		return response.ErrCodeInternalServerError, err
	}

	// Insert Info Account To User Profile MySQL
	result, err := ua.repo.InsertUserProfileRegister(context.Background(), database.InsertUserProfileRegisterParams{
		UserID: int32(lastIdUserBase),
		UserEmail: infoVerify.VerifyKey,
	})
	if err != nil {
		return response.ErrCodeInternalServerError, err
	}
	
	lastIdUser, err := result.LastInsertId()
	if err != nil {
		return response.ErrCodeInternalServerError, err
	}
	global.Logger.Info("User Register Full", zap.Int("user_id", int(lastIdUser)))
	return response.CodeSuccess, nil
}

func (ua *userAuthImpl) ForgotPassword(in *model.ForgotPasswordRequest) (int, error) {
	return response.CodeSuccess, nil
}


func (ua *userAuthImpl) Logout() (int, error) {
	return response.CodeSuccess, nil
}

func NewUserAuthImpl(repo *database.Queries) *userAuthImpl {
	return &userAuthImpl{
		repo: repo,
	}
}

