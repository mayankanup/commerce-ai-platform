package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/config"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/logging"
	"github.com/mayankanup/commerce-ai-platform/internal/platform/server"
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
	/*logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)*/

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "commerce-ai-platform",
			"status":  "running",
			"version": "0.1.0",
		})
	})

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	server := server.New(server.Options{
		Address: address,
		Logger:  logger,
		Handler: router,
	})

	go func() {
		slog.Info(
			"Configuration loaded",
			"app", cfg.App.Name,
			"version", cfg.App.Version,
			"environment", cfg.App.Environment,
			"port", cfg.Server.Port,
			"model", cfg.Ollama.Model,
		)

		if err := server.Start(); err != nil {
			logger.Error(
				"Server stopped",
				"error",
				err,
			)
			os.Exit(1)
		}
		/*if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed", "error", err)
			os.Exit(1)
		}*/
	}()

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
