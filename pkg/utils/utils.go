package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)


var validate = validator.New()


func GetUserKey(key string) string {
	return fmt.Sprintf("u:%s:otp", key)
}


func ParseJSON(r *http.Request, dest interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dest)
	if err != nil {
		return err
	}

	err = validate.Struct(dest)

	if err != nil {
		return err
	}

	return nil
}