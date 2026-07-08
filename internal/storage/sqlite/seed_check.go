package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

func (d *Database) isSeeded(ctx context.Context) (bool, error) {

	var count int

	err := d.db.QueryRowContext(
		ctx,
		`SELECT COUNT(*) FROM categories`,
	).Scan(&count)

	if err != nil {

		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, fmt.Errorf("check seed status: %w", err)
	}

	return count > 0, nil
}
