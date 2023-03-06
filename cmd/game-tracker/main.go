package main

import (
	"flag"
	"fmt"
	"github.com/HeadGardener/game-tracker/configs"
	"github.com/HeadGardener/game-tracker/internal/app/handlers"
	"github.com/HeadGardener/game-tracker/internal/app/repositories"
	"github.com/HeadGardener/game-tracker/internal/app/services"
	"github.com/HeadGardener/game-tracker/internal/pkg/server"
	"go.uber.org/zap"
)

var confPath = flag.String("conf-path", "./configs/.env", "path to config env")

func main() {
	logger, _ := zap.NewProduction()

	dbconfig, err := configs.NewDBConfig(*confPath)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to read config file, error:%d", err))
	}

	dbConn, err := repositories.NewDBConn(*dbconfig)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to make up conn with db, error:%d", err))
	}

	repository := repositories.NewRepository(dbConn)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)

	srvconfig, err := configs.NewServerConfig(*confPath)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to read config file, error:%d", err))
	}
	srv := server.New(*logger)
	if err := srv.Run(srvconfig.ServerPort, handler.InitRoutes()); err != nil {
		// change later
		logger.Fatal(fmt.Sprintf("error occurring while running server, err:%d", err))
	}
}
