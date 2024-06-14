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
	"github.com/go-redis/redis"
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
	r := redis.NewClient(&redis.Options{
		Addr:     "redis-13806.c273.us-east-1-2.ec2.redns.redis-cloud.com:13806",
		Password: "ONxog6dHNbz00ycjdHFk75GxHDrJ3Ep6",
		DB:       0,
	})
	srv := service.GetService(repos, r)
	handlers := handler.GetHandler(srv, zLogger)
	routes := route.GetRoute(handlers)
	err = server.StartListen(routes, &conf.Srv)
	if err != nil {
		zLogger.Error("Can't start listen and serve", zap.Error(err))
		return
	}
}
