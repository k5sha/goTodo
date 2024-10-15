package main

import (
	"github.com/k5sha/golang-todo-example/internal/config"
	httpServer "github.com/k5sha/golang-todo-example/internal/http-server"
	"github.com/k5sha/golang-todo-example/internal/lib/logger/sl"
	"github.com/k5sha/golang-todo-example/internal/storage/postgresql"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting todoModels-app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := postgresql.New(cfg.Database, log)
	if err != nil {
		log.Error("failed to connect db", sl.Err(err))
		os.Exit(1)
	}

	router := httpServer.SetupRoutes(log, storage)

	//	Startup server

	log.Info("starting server", slog.String("address", cfg.Address))

	server := httpServer.NewServer(
		cfg.Address,
		router,
		cfg.HttpServer.Timeout,
		cfg.HttpServer.IdleTimeout,
	)

	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start server", sl.Err(err))
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
