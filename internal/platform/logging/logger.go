package logging

import (
	"io"
	"log/slog"
	"os"
)

type Logger = slog.Logger

func New(options Options) *slog.Logger {
	writer := options.Output
	if writer == nil {
		writer = os.Stdout
	}
	return NewWithWriter(options, writer)
}

// NewWithWriter creates a logger that writes to the provided writer.
// This is primarily used by unit tests.
func NewWithWriter(options Options, writer io.Writer) *slog.Logger {

	handlerOptions := &slog.HandlerOptions{
		Level: options.Level,
	}

	var handler slog.Handler

	switch options.Format {

	case "text":
		handler = slog.NewTextHandler(
			writer,
			handlerOptions,
		)

	case "json":
		fallthrough

	default:
		handler = slog.NewJSONHandler(
			writer,
			handlerOptions,
		)
	}

	return slog.New(handler)
}
