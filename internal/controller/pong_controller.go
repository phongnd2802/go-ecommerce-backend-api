package controller

import (
	"net/http"
	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/response"
)

// ShowAccount godoc
// @Summary      Pong
// @Description  Monitor Check
// @Tags         Health Check
// @Accept       json
// @Produce      json
// @Success      200
// @Router       / [get]
func Pong(w http.ResponseWriter, r *http.Request) {
	response.SuccessResponse(w, response.CodeSuccess, "OK!")
}