package main

import (
	"context"
	"fmt"
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

	// Load configuration
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		panic(err)
	}

	// Create logger
	logger := logging.New(
		logging.Options{
			Level:  logging.ParseLevel(cfg.Logging.Level),
			Format: cfg.Logging.Format,
		},
	)

	logger.Info(
		"Commerce AI Platform starting",
		"version", cfg.App.Version,
		"environment", cfg.App.Environment,
	)

	// Create router
	router := web.NewRouter(
		web.ApplicationInfo{
			Name:        cfg.App.Name,
			Version:     cfg.App.Version,
			Environment: cfg.App.Environment,
		},
		logger,
	)

	// Create HTTP server
	srv := server.New(
		server.Options{
			Address: fmt.Sprintf(
				"%s:%d",
				cfg.Server.Host,
				cfg.Server.Port,
			),
			Logger:  logger,
			Handler: router,
		},
	)

	// Start server in background
	go func() {
		if err := srv.Start(); err != nil {
			logger.Error(
				"HTTP server stopped unexpectedly",
				"error", err,
			)
			os.Exit(1)
		}
	}()

	// Wait for termination signal
	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		os.Interrupt,
		syscall.SIGTERM,
	)

	sig := <-signalChan

	logger.Info(
		"Shutdown signal received",
		"signal", sig.String(),
	)

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(
			"Failed to gracefully shutdown server",
			"error", err,
		)
		os.Exit(1)
	}

	logger.Info("Commerce AI Platform stopped")
}
