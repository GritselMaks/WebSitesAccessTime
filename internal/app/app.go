package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"siteavliable/internal/configs"
	adminhandler "siteavliable/internal/controllers/admin_handler"
	clienthandler "siteavliable/internal/controllers/client_handler"
	"siteavliable/internal/metrics"
	statscase "siteavliable/internal/usecases/clientsCase"
	"siteavliable/internal/usecases/repo"
	servicese "siteavliable/internal/usecases/statsCase"
	updatecase "siteavliable/internal/usecases/updateCase"
	redisclient "siteavliable/pkg/client/redis"
	httpserver "siteavliable/pkg/httpserver"
	"siteavliable/pkg/utils"
	"time"
)

type app struct {
	cfg         *configs.Config
	logger      *log.Logger
	redisClient *redisclient.RedisClient
}

// New returns a new app instance and an error
func New(cfg *configs.Config, l *log.Logger) (*app, error) {
	app := app{
		cfg:    cfg,
		logger: l,
	}
	return &app, nil
}

func (a *app) configStore(ctx context.Context) error {
	client, err := redisclient.NewRedisClient(ctx, a.cfg.Redis)
	if err != nil {
		a.logger.Panicf("error conect to Redis. Error: %s", err.Error())
		return err
	}
	a.redisClient = client
	return nil
}

// Run runs app.
// It's blocks until the http.server runs
func (a *app) Run(ctx context.Context) error {
	err := a.configStore(ctx)
	if err != nil {
		a.logger.Printf("ErrorConfigStore. Error: %s\n", err.Error())
		return err
	}
	urls, err := utils.LoadUrlsList(a.cfg.FilePath)
	if err != nil {
		a.logger.Printf("ErrorRead Websites list.  Error: %s\n", err.Error())
		return err
	}
	//configure controllers
	urlRepo := repo.NewUrlsRepo(a.redisClient, a.cfg.Redis.Set)
	statsRepo := repo.NewSatsRepo(a.redisClient)
	clientUCase := statscase.New(urlRepo, a.logger)
	statUCase := servicese.New(statsRepo, a.logger, []string{metrics.GetWithMaxTime, metrics.GetWithMinTime, metrics.GetByURL})
	mux := http.NewServeMux()
	clienthandler.ClientsRouter(mux, clientUCase, a.logger)
	adminhandler.AdminRouter(mux, statUCase, a.logger, a.cfg.AdminAuth)

	// Run func for update access time
	go func() {
		updater := updatecase.New(urlRepo, a.logger, urls)
		ticker := time.NewTicker(time.Second * time.Duration(a.cfg.UpdateTimeout))
		updater.UpdateAccessTime(ctx)
		for {
			select {
			case <-ticker.C:
				ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration((a.cfg.UpdateTimeout-3)))
				defer cancel()
				updater.UpdateAccessTime(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()

	// Run func for save metrics
	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(a.cfg.SaveMetricsTimeout))
		for {
			select {
			case <-ticker.C:
				statUCase.SaveMetrics(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()

	//run http Server
	srv, err := httpserver.New(a.cfg.Port)
	if err != nil {
		a.logger.Printf("server.New Error: %s\n", err.Error())
		return fmt.Errorf("server.New: %w", err)
	}
	a.logger.Printf("server starting on port %s.....\n", a.cfg.Port)
	return srv.ServeHTTPHandler(ctx, mux)
}
