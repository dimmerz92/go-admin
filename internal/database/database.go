package database

import (
	"context"
	"database/sql"
	"fmt"
	"maps"
	"slices"

	"github.com/dimmerz92/go-admin/internal/core"
	"github.com/pressly/goose/v3"
)

var validDrivers = map[string]goose.Dialect{
	"sqlite": goose.DialectSQLite3,
}

var drivers = slices.Sorted(maps.Keys(validDrivers))

// OpenDB returns a sql database connection with a fully managed lifecycle.
// Cancelling the context will automatically close the connection.
func OpenDB(ctx context.Context, driver, dsn string) (*sql.DB, error) {
	if _, ok := validDrivers[driver]; !ok {
		return nil, fmt.Errorf("%w: unsupported driver %s: supported drivers %v", core.ErrDatabase, driver, drivers)
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", core.ErrDatabase, err)
	}

	if err = db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("%w: %w", core.ErrDatabase, err)
	}

	go func() { <-ctx.Done(); _ = db.Close() }()

	return db, nil
}
