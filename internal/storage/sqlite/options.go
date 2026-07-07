package sqlite

import (
	"log/slog"
	"time"
)

type Options struct {
	Path            string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	Logger          *slog.Logger
}
