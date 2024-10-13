package router

import (
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/router/admin"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/router/owner"
	"github.com/phongnd2802/go-ecommerce-backend-api/internal/router/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Owner owner.OwnerRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterApp = new(RouterGroup)
