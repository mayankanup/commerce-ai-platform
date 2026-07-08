package sqlite

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

func (d *Database) Seed(ctx context.Context) error {

	seeded, err := d.isSeeded(ctx)
	if err != nil {
		return err
	}

	if seeded {
		d.logger.Info("Database already contains seed data. Skipping.")
		return nil
	}

	files, err := d.seedFiles()
	if err != nil {
		return err
	}

	for _, file := range files {

		d.logger.Info("Executing seed file",
			"file", filepath.Base(file),
		)

		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("read seed file %s: %w", file, err)
		}

		if _, err := d.db.ExecContext(ctx, string(sqlBytes)); err != nil {
			return fmt.Errorf("execute seed file %s: %w", file, err)
		}
	}

	d.logger.Info("Database seeded successfully")

	return nil
}
