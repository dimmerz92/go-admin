package sqlite

import (
	"context"

	"github.com/dimmerz92/go-admin/configs"
	"modernc.org/sqlite"
)

const DRIVER configs.DatabaseDriver = "sqlite"

var pragmas = []string{
	"PRAGMA busy_timeout = 5000",
	"PRAGMA foreign_keys = ON",
	"PRAGMA journal_mode = WAL",
	"PRAGMA synchronous = FULL",
}

func init() {
	sqlite.RegisterConnectionHook(func(conn sqlite.ExecQuerierContext, dsn string) error {
		for _, pragma := range pragmas {
			if _, err := conn.ExecContext(context.Background(), pragma, nil); err != nil {
				return err
			}
		}
		return nil
	})
}
