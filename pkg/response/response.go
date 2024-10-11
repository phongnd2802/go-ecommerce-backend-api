package response

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func SuccessResponse(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ResponseData{
		Code: code,
		Message: msg[code],
		Data: data,
	})
}


func ErrorResponse(w http.ResponseWriter, code int, errStr string) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(ResponseData{
		Code: code,
		Message: msg[code],
		Data: errStr,
	})
}

