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

// @title account api
// @version 1.0
// @description Seed-IOT

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	go func() {
		r := gin.New()
		url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		r.Run()
	}()

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

	server := web.NewServer(cfg, service)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.GlobalLog.WithFields(logrus.Fields{
				"error": err,
			}).Info("Failed to listen for http server")
		}
	}()

	// 程序结束 关闭db
	defer service.DB.Close()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	log.GlobalLog.Info("account is running")
	<-quit
	log.GlobalLog.Info("account is stopped")

}
