package sqlite

import (
	"context"
	"fmt"
	"os"
)

func (d *Database) Migrate(ctx context.Context, schemaPath string) error {

	d.logger.Info("Running database migrations",
		"schema", schemaPath,
	)

	sqlBytes, err := os.ReadFile(schemaPath)
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
