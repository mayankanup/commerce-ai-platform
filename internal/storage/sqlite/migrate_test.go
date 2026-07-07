package sqlite

import (
	"context"
	"path/filepath"
	"testing"
)

func TestMigrate(t *testing.T) {
	db, err := New(
		Options{
			Path: filepath.Join(
				t.TempDir(),
				"test.db",
			),
			SchemaPath: "../../../data/schema.sql",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	err = db.Migrate(
		context.Background(),
	)

	if err != nil {
		t.Fatal(err)
	}

	var version int

	if err := db.DB().QueryRow(
		"SELECT version FROM schema_version",
	).Scan(&version); err != nil {
		t.Fatal(err)
	}

	if version != 1 {
		t.Fatalf("expected schema version 1, got %d", version)
	}
}
