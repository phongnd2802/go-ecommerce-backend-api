package router

import "github.com/phongnd2802/go-ecommerce-backend-api/internal/router/user"


type RouterGroup struct {
	User user.UserRouterGroup
}

var RouterApp = new(RouterGroup)
