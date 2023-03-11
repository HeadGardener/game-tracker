package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/HeadGardener/game-tracker/configs"
	"github.com/HeadGardener/game-tracker/internal/app/handlers"
	"github.com/HeadGardener/game-tracker/internal/app/repositories"
	"github.com/HeadGardener/game-tracker/internal/app/services"
	"github.com/HeadGardener/game-tracker/internal/pkg/server"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
	"time"
)

var confPath = flag.String("conf-path", "./configs/.env", "path to config env")

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger, _ := zap.NewProduction()

	dbconfig, err := configs.NewDBConfig(*confPath)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to read config file, error:%e", err))
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

	srv := &server.Server{}

	go func() {
		if err := srv.Run(srvconfig.ServerPort, handler.InitRoutes()); err != nil {
			logger.Error(fmt.Sprintf("error occurring while running server, err:%d", err))
		}
	}()

	logger.Info("server start working")
	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("server forced to shutdown: %d", err))
	}

	if err := dbConn.Close(ctx); err != nil {
		logger.Error(fmt.Sprintf("db connection forced to shutdown: %d", err))
	}

	logger.Info("server exiting")
}
