package main

import (
	"back-end/internal/config"
	"back-end/internal/db"
	"back-end/internal/handler"
	"back-end/internal/repository"
	"back-end/internal/service"
	"back-end/pkg/logger"
	"go.uber.org/zap"
	"log"
)

func main() {
	zLogger, err := logger.GetLogger()
	if err != nil {
		log.Println(err)
		return
	}
	conf, err := config.GetConfigs()
	if err != nil {
		zLogger.Error("Can't get configs.", zap.Error(err))
		return
	}
	db, err := db.GetConnection(&conf.Db)
	if err != nil {
		zLogger.Error("Can't get connection.", zap.Error(err))
		return
	}
	repos := repository.GetRepository(db)
	srv := service.GetService(repos)
	handler := handler.GetHandler(srv)

}
