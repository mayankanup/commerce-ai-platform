package server

import (
	"log/slog"
	"net/http"
	"time"
)

type Options struct {
	Address string
	Logger  *slog.Logger
	Handler http.Handler

	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}
