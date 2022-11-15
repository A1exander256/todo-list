package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/a1exander256/todo/config"
	"github.com/a1exander256/todo/logger"
	"github.com/a1exander256/todo/pkg/handler"
	"github.com/a1exander256/todo/pkg/service"
	"github.com/a1exander256/todo/pkg/storage"
	"github.com/a1exander256/todo/server"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	conf = kingpin.Flag("conf", "path to the config file").Short('c').Default("/app/config/config-docker.json").String()
)

func main() {
	kingpin.Parse()

	config, err := config.InitConfig(*conf)
	if err != nil {
		log.Fatal(err)
	}

	log := logger.InitLogger(config.Logger.Level)

	db, err := storage.NewPostgresDB(&config.Postgres)
	if err != nil {
		log.Fatal("failed to initialize db : ", err)
	}
	storage := storage.NewStorage(db, log)
	service := service.NewService(storage, log)
	handler := handler.NewHandler(service, log)

	server := server.NewServer()

	go func() {
		if err := server.Run(config.Gin.Port, handler.InitRoutes(config.Gin.Mode)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Info("app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("app shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
