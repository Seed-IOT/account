package main

import (
	"account/config"
	"account/log"
	"account/model"
	"account/web"

	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
)

func main() {

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
