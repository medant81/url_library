package main

import (
	"context"
	"github.com/medant81/url_library/internal/author"
	"github.com/medant81/url_library/internal/book"
	"github.com/medant81/url_library/internal/config"
	"github.com/medant81/url_library/internal/handlers"
	"github.com/medant81/url_library/internal/storage/postgre"
	"github.com/medant81/url_library/version"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// @title url_library App API
// @version 1.0
// @description API Server for url_library Application

// @host localhost:3000
// @BasePath /
func main() {

	os.Setenv("CONFIG_PATH", "./config/local.yaml")

	cfg := config.MustLoad()

	// TODO: init logger: log/slog
	log := setupLogger(cfg.Env)

	log.Info(
		"Start url library",
		slog.String("env", cfg.Env),
		slog.String("version", version.Version()),
	)
	log.Debug("Debug messages a enabled")
	log.Debug("env: ", cfg)

	// TODO: init storage: postgre
	clientPSQL, err := postgre.NewClient(context.Background(), 3, cfg.StorageConfig, log)
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}
	log.Info("Start clientPSQL")

	// TODO: init router: net/http
	var app handlers.Application
	app = handlers.Application{
		RBook:   book.NewRepository(clientPSQL, log),
		RAuthor: author.NewRepository(clientPSQL, log),
	}
	log.Info("Init handlers application")

	// TODO: run server
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      handlers.Routers(&app),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	log.Debug("Start http server: ", srv)

	err = srv.ListenAndServe()
	if err != nil {
		log.Error("connection error on 8080!", err)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
