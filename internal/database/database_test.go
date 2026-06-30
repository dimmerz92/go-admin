package database_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/dimmerz92/go-admin/configs/drivers/sqlite"
	"github.com/dimmerz92/go-admin/internal/database"
)

func TestOpen(t *testing.T) {
	tests := []struct {
		name        string
		driver      string
		dsn         string
		expectedErr bool
	}{
		{name: "valid open", driver: sqlite.DRIVER, dsn: ":memory:", expectedErr: false},
		{name: "invalid driver", driver: "invalid", dsn: ":memory:", expectedErr: true},
		{name: "invalid dsn", driver: sqlite.DRIVER, dsn: "file:invalid?mode=ro", expectedErr: true},
	}

	expected := 4321

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(t.Context())
			defer cancel()

			db, err := database.OpenDB(ctx, test.driver, test.dsn)
			if test.expectedErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("failed to open database: %v", err)
			}

			var got int
			err = db.QueryRowContext(ctx, fmt.Sprintf("SELECT %d", expected)).Scan(&got)
			if err != nil {
				t.Fatalf("failed to query database: %v", err)
			}

			if got != expected {
				t.Fatalf("expected value %d, got %d", expected, got)
			}
		})
	}
}
