package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}


func HashPasswordWithSalt(password string, salt string) (string, error) {
	passwordWithSalt := password + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordWithSalt(password, salt, hashedPassword string) bool {
	passwordWithSalt := password + salt
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordWithSalt))
	return err == nil
}