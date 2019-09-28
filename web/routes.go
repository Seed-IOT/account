package web

import (
	"account/middleware"
	"net/http"
)

func (srv *server) routes() http.Handler {

	md := &middleware.Middleware{}

	srv.router.POST("account/mobileLogin", srv.MobileLogin)
	srv.router.POST("account/login", srv.Login)
	srv.router.POST("account/getCode", srv.GetCode)
	srv.router.Use(md.Authorize())
	srv.router.GET("account/version/:name", srv.Version)

	//Declare web routing table at here.
	return srv.router
}
