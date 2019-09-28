package main

import (
	"account/config"
	"account/log"
	"account/model"
	"account/web"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "account/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description kylewang.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := gin.New()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()

	log.Init()

	cfg, err := config.New()
	if err != nil {
		log.GlobalLog.WithFields(logrus.Fields{
			"error": err,
		}).Info("Failed to reading config file")
	}

	service, err := model.New(cfg.Database)
	if err != nil {
		log.GlobalLog.WithFields(logrus.Fields{
			"error": err,
		}).Info("Failed to initialize model for operating all service")
	}

	// 程序结束 关闭db
	defer service.DB.Close()

	server := web.NewServer(cfg, service)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.GlobalLog.WithFields(logrus.Fields{
				"error": err,
			}).Info("Failed to listen for http server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	log.GlobalLog.Info("account is running")
	<-quit
	log.GlobalLog.Info("account is stopped")
}
