package sqlite

import (
	"context"
)

func (d *Database) Health(ctx context.Context) error {
	return d.db.PingContext(ctx)
}
