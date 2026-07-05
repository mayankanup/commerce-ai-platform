package sqlite

import (
	"context"
	"path/filepath"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	db, err := New(
		Options{
			Path: ":memory:",
		},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer db.Close()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	if err := db.Health(ctx); err != nil {
		t.Fatalf("health check failed: %v", err)
	}
}

func TestNewEmptyPath(t *testing.T) {

	_, err := New(
		Options{},
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNewCreatesParentDirectory(t *testing.T) {
	path := filepath.Join(
		t.TempDir(),
		"data",
		"ecommerce.db",
	)

	db, err := New(
		Options{
			Path: path,
		},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer db.Close()
}
