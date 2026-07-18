package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	logger *slog.Logger
	server *http.Server
}

func New(options Options) *Server {

	if options.Logger == nil {
		options.Logger = slog.Default()
	}

	if options.ReadTimeout == 0 {
		options.ReadTimeout = 30 * time.Second
	}

	if options.WriteTimeout == 0 {
		options.WriteTimeout = 5 * time.Minute
	}

	if options.IdleTimeout == 0 {
		options.IdleTimeout = 2 * time.Minute
	}

	httpServer := &http.Server{
		Addr:         options.Address,
		Handler:      options.Handler,
		ReadTimeout:  options.ReadTimeout,
		WriteTimeout: options.WriteTimeout,
		IdleTimeout:  options.IdleTimeout,
	}

	return &Server{
		logger: options.Logger,
		server: httpServer,
	}
}

func (s *Server) Start() error {

	s.logger.Info(
		"Starting HTTP server",
		"address",
		s.server.Addr,
	)

	err := s.server.ListenAndServe()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {

	s.logger.Info("Shutting down HTTP server")

	return s.server.Shutdown(ctx)
}
