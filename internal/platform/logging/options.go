package logging

import (
	"io"
	"log/slog"
)

type Options struct {
	Level       slog.Level
	Format      string
	Output      io.Writer
	ServiceName string
	Version     string
	Environment string
}
