package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mixdjoker/auth/internal/client/db"
	"github.com/pkg/errors"
)

const (
	newErr = "pgxpool.New"
)

type pgClient struct {
	masterDBC db.DB
}

// NewClient creates a new database client.
func NewClient(ctx context.Context, dsn string) (db.Client, error) {
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, errors.Wrap(err, newErr)
	}

	pgDB := NewPool(dbpool)

	return &pgClient{masterDBC: pgDB}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
