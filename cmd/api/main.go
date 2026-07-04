package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/logging"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/server"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/web"
)

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}
	logger := logging.New(logging.Options{
		Level:  logging.ParseLevel(cfg.Logging.Level),
		Format: cfg.Logging.Format,
	})
	logger.Info(
		"Commerce AI Platform starting",
		"version", cfg.App.Version,
		"environment", cfg.App.Environment,
		"port", cfg.Server.Port,
	)

	router := web.NewRouter(
		web.ApplicationInfo{
			Name:        cfg.App.Name,
			Version:     cfg.App.Version,
			Environment: cfg.App.Environment,
		},
		logger,
	)

	server := server.New(
		server.Options{
			Address: fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
			Logger:  logger,
			Handler: router,
		},
	)

	if err := server.Start(); err != nil {
		logger.Error(
			"server stopped",
			"error",
			err,
		)
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	slog.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Graceful shutdown failed", "error", err)
	}

	slog.Info("Server stopped")
}
