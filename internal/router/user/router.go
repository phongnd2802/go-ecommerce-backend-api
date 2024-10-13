package user

import "github.com/gorilla/mux"

type UserRouterGroup struct {
	UserRouter
}

func (r *UserRouterGroup) InitRouter(Router *mux.Router) {
	r.initUserRouter(Router)
}