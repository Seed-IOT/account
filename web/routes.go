package web

import (
	"account/middleware"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (srv *server) routes() http.Handler {

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	srv.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	md := &middleware.Middleware{}

	srv.router.POST("account/mobileLogin", srv.MobileLogin)
	srv.router.POST("account/login", srv.Login)
	srv.router.POST("account/getCode", srv.GetCode)
	srv.router.Use(md.Authorize())
	srv.router.GET("account/version/:name", srv.Version)

	//Declare web routing table at here.
	return srv.router
}
