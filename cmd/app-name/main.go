package main

import (
	"back-end/internal/config"
	"back-end/internal/handler"
	"back-end/internal/repository"
	"back-end/internal/route"
	"back-end/internal/server"
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
	//db, err := db.GetConnection(&conf.Db)
	//if err != nil {
	//	zLogger.Error("Can't get connection.", zap.Error(err))
	//	return
	//}
	repos := repository.GetRepository()
	srv := service.GetService(repos)
	handlers := handler.GetHandler(srv)
	routes := route.GetRoute(handlers)
	err = server.StartListen(routes, &conf.Srv)
	if err != nil {
		zLogger.Error("Can't start listen and serve", zap.Error(err))
	}
}
