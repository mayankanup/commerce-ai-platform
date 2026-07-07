package sqlite

import (
	"context"
	"fmt"
	"os"
)

func (d *Database) Migrate(ctx context.Context) error {
	d.logger.Info(
		"Running database migrations",
		"schemaPath", d.options.SchemaPath,
	)

	sqlBytes, err := os.ReadFile(d.options.SchemaPath)
	if err != nil {
		return fmt.Errorf("read schema: %w", err)
	}

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, string(sqlBytes)); err != nil {
		return fmt.Errorf("execute schema: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit migration: %w", err)
	}

	d.logger.Info("Database migration completed")

	return nil
}
