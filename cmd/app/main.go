package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"siteavliable/internal/app"
	"siteavliable/internal/configs"
	"siteavliable/internal/metrics"
	"syscall"
)

var configPath = flag.String("config", "./app.env", "service config")

func main() {
	// Loading the config
	flag.Parse()
	cfg, err := configs.LoadConfig(*configPath)
	if err != nil {
		log.Fatal("App::Initialize load config error: ", err)
	}
	// Сonfigure the logger
	file, err := os.OpenFile(cfg.LogPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("error Logger: ", err.Error())
	}
	defer file.Close()
	logger := log.New(file, "accesss time app ", log.LstdFlags|log.Lshortfile)

	// Init metrics
	metrics.Init()

	// Processing a stop signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		stop()
		if r := recover(); r != nil {
			logger.Print("application panic", "panic", r)
			file.Close()
			os.Exit(1)
		}
	}()

	err = realMain(ctx, cfg, logger)
	stop()

	if err != nil {
		logger.Fatal("error server run: ", err.Error())
	}

	logger.Println("successful shutdown")
}

func realMain(ctx context.Context, cfg *configs.Config, l *log.Logger) error {
	app, err := app.New(cfg, l)
	if err != nil {
		return fmt.Errorf("error new app: %w", err)
	}
	return app.Run(ctx)
}
