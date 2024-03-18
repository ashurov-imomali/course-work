package main

import (
	"back-end/internal/config"
	"back-end/internal/db"
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
	defer zLogger.Sync()
	conf, err := config.GetConfigs()
	if err != nil {
		zLogger.Error("Can't get configs.", zap.Error(err))
		return
	}
	connection, err := db.GetConnection(&conf.Db)
	if err != nil {
		zLogger.Error("Can't get connection.", zap.Error(err))
		return
	}
	repos := repository.GetRepository(connection)
	srv := service.GetService(repos)
	handlers := handler.GetHandler(srv, zLogger)
	routes := route.GetRoute(handlers)
	err = server.StartListen(routes, &conf.Srv)
	if err != nil {
		zLogger.Error("Can't start listen and serve", zap.Error(err))
	}
}
