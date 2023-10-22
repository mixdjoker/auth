package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mixdjoker/auth/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

func NewClient(ctx context.Context, dsn string) (db.Client, error) {
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
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
