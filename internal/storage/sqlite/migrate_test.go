package sqlite

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

func TestMigrate(t *testing.T) {

	tmp, err := os.CreateTemp("", "*.db")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmp.Name())

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := sqlite.New(
		sqlite.Config{
			Path: tmp.Name(),
		},
		logger,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	err = db.Migrate(
		context.Background(),
		"../../../data/schema.sql",
	)

	if err != nil {
		t.Fatal(err)
	}
}
