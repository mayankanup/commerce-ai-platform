package logging

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestJSONLogger(t *testing.T) {

	var buf bytes.Buffer

	logger := NewWithWriter(
		Options{
			Level:  slog.LevelInfo,
			Format: "json",
		},
		&buf,
	)

	logger.Info("hello world")

	output := buf.String()

	if !strings.Contains(output, "hello world") {
		t.Fatalf("expected log message to contain 'hello world'")
	}

	if !strings.Contains(output, "\"level\":\"INFO\"") {
		t.Fatalf("expected INFO level in JSON log")
	}
}

func TestTextLogger(t *testing.T) {

	var buf bytes.Buffer

	logger := NewWithWriter(
		Options{
			Level:  slog.LevelInfo,
			Format: "text",
		},
		&buf,
	)

	logger.Info("hello world")

	output := buf.String()

	if !strings.Contains(output, "hello world") {
		t.Fatalf("expected log message to contain 'hello world'")
	}

	if !strings.Contains(output, "INFO") {
		t.Fatalf("expected INFO level in text log")
	}
}

func TestDebugLoggingEnabled(t *testing.T) {

	var buf bytes.Buffer

	logger := NewWithWriter(
		Options{
			Level:  slog.LevelDebug,
			Format: "json",
		},
		&buf,
	)

	logger.Debug("debug message")

	output := buf.String()

	if !strings.Contains(output, "debug message") {
		t.Fatalf("expected debug message to be logged")
	}
}

func TestDebugLoggingDisabled(t *testing.T) {

	var buf bytes.Buffer

	logger := NewWithWriter(
		Options{
			Level:  slog.LevelInfo,
			Format: "json",
		},
		&buf,
	)

	logger.Debug("debug message")

	output := buf.String()

	if output != "" {
		t.Fatalf("expected debug message to be filtered")
	}
}

func TestInvalidFormatDefaultsToJSON(t *testing.T) {

	var buf bytes.Buffer

	logger := NewWithWriter(
		Options{
			Level:  slog.LevelInfo,
			Format: "invalid",
		},
		&buf,
	)

	logger.Info("hello")

	output := buf.String()

	if !strings.Contains(output, "\"level\":\"INFO\"") {
		t.Fatalf("expected JSON logger for invalid format")
	}
}

func TestParseLevel(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected slog.Level
	}{
		{"debug", "debug", slog.LevelDebug},
		{"info", "info", slog.LevelInfo},
		{"warn", "warn", slog.LevelWarn},
		{"error", "error", slog.LevelError},
		{"uppercase", "DEBUG", slog.LevelDebug},
		{"spaces", " info ", slog.LevelInfo},
		{"invalid", "foobar", slog.LevelInfo},
		{"empty", "", slog.LevelInfo},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			level := ParseLevel(tt.input)

			if level != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, level)
			}
		})
	}
}
