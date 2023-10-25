package pg

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mixdjoker/auth/internal/client/db"
	prettier "github.com/mixdjoker/auth/internal/client/db/prettifier"
)

type key string

const (
	// TxKey is a key for transaction.
	TxKey key = "tx"
)

// pg is a database client wrapper for pgx.Pool.
type pg struct {
	dbc *pgxpool.Pool
}

// NewPool creates a new database client.
func NewPool(dbc *pgxpool.Pool) *pg {
	return &pg{dbc: dbc}
}

// ScanOneContext scans one row from database.
func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	return pgxscan.ScanOne(dest, rows)
}

// ScanAllContext scans all rows from database.
func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	return pgxscan.ScanAll(dest, rows)
}

// ExecContext executes a query without returning any rows.
func (p *pg) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	if tx, ok := ctx.Value(TxKey).(pgx.Tx); ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT.
func (p *pg) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	if tx, ok := ctx.Value(TxKey).(pgx.Tx); ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext executes a query that is expected to return at most one row.
func (p *pg) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	if tx, ok := ctx.Value(TxKey).(pgx.Tx); ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

// BeginTx begins a transaction.
func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

// Close closes the database client.
func (p *pg) Close() {
	p.dbc.Close()
}

// MakeContextTx makes a new context with transaction.
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func logQuery(ctx context.Context, q db.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)

	buf := strings.Builder{}
	fmt.Fprint(&buf, "query: {\n")
	fmt.Fprintf(&buf, "\tq.ctx: %+v,\n", ctx)
	fmt.Fprintf(&buf, "\tq.Name: %s,\n", q.Name)
	fmt.Fprintf(&buf, "\tq.SQL: %s\n", prettyQuery)
	fmt.Fprint(&buf, "}")

	log.Println(color.MagentaString(buf.String()))
}
