package jwt

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(payload map[string]interface{}, secretKey *rsa.PrivateKey, expirationTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": payload,
		"exp": time.Now().Add(time.Duration(expirationTime) * time.Second).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func VerifyToken(tokenStr string, publicKeyStr string) (*jwt.Token, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyStr))
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func GetPayloadToken(token *jwt.Token) (map[string]interface{}, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if payload, ok := claims["sub"].(map[string]interface{}); ok {
		return payload, nil
	}

	return nil, fmt.Errorf("payload invalid type")
}