package logging

import (
	"log/slog"
	"strings"
)

func ParseLevel(level string) slog.Level {

	switch strings.ToLower(level) {

	case "debug":
		return slog.LevelDebug

	case "warn":
		return slog.LevelWarn

	case "error":
		return slog.LevelError

	default:
		return slog.LevelInfo
	}
}
