package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

type Database struct {
	db      *sql.DB
	logger  *slog.Logger
	options Options
}

func New(options Options) (*Database, error) {

	if options.Path == "" {
		return nil, fmt.Errorf("database path is required")
	}

	if options.SchemaPath == "" {
		return nil, fmt.Errorf("schema path is required")
	}

	if err := ensureParentDir(options.Path); err != nil {
		return nil, err
	}

	logger := options.Logger
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}

	db, err := sql.Open(
		"sqlite",
		options.Path,
	)
	if err != nil {
		return nil, err
	}

	if options.MaxOpenConns > 0 {
		db.SetMaxOpenConns(options.MaxOpenConns)
	}

	if options.MaxIdleConns > 0 {
		db.SetMaxIdleConns(options.MaxIdleConns)
	}

	if options.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(options.ConnMaxLifetime)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return &Database{
		db:      db,
		logger:  logger,
		options: options,
	}, nil
}

func (d *Database) DB() *sql.DB {
	return d.db
}

func (d *Database) Close() error {
	return d.db.Close()
}

func ensureParentDir(path string) error {
	if path == ":memory:" || strings.HasPrefix(path, "file:") {
		return nil
	}

	dir := filepath.Dir(path)
	if dir == "." {
		return nil
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("unable to create database directory %q: %w", dir, err)
	}

	return nil
}
