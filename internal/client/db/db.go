package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Handler is a type of functions that is executed in transaction.
type Handler func(ctx context.Context) error

// Client is a database client.
type Client interface {
	DB() DB
	Close() error
}

// DB is an interface for database driver.
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}

// SQLExecer complex NamedExecer & QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer is an interface for database named query.
type NamedExecer interface {
	ScanOneContext(context.Context, interface{}, Query, ...interface{}) error
	ScanAllContext(context.Context, interface{}, Query, ...interface{}) error
}

// QueryExecer is an interface for database common query.
type QueryExecer interface {
	ExecContext(context.Context, Query, ...interface{}) (pgconn.CommandTag, error)
	QueryContext(context.Context, Query, ...interface{}) (pgx.Rows, error)
	QueryRowContext(context.Context, Query, ...interface{}) pgx.Row
}

// Transactor is an interface for beggin database transaction.
type Transactor interface {
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

// Pinger is an interface for database ping.
type Pinger interface {
	Ping(context.Context) error
}

// Query is a named query struct.
type Query struct {
	Name     string
	QueryRaw string
}

// TxManager is an interface for transaction manager,
type TxManager interface {
	ReadCommitted(context.Context, Handler) error
}
