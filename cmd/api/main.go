package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mayankanup/commerce-ai-platform/internal/app"
)

func main() {
	application, err := app.Bootstrap(
		app.Options{
			ConfigFile: "config/config.yaml",
		},
	)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := application.Start(); err != nil {
			application.Logger.Error(
				"Application stopped unexpectedly",
				"error", err,
			)
			os.Exit(1)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		os.Interrupt,
		syscall.SIGTERM,
	)

	sig := <-signalChan

	application.Logger.Info(
		"Shutdown signal received",
		"signal", sig.String(),
	)

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := application.Shutdown(ctx); err != nil {
		application.Logger.Error(
			"Application shutdown failed",
			"error", err,
		)
		os.Exit(1)
	}

	application.Logger.Info("Commerce AI Platform stopped")
}
