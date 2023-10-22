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
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer is an interface for database common query.
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// TxManager is a transaction manager which provides a user-specified transaction handler
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Pinger is an interface for database ping.
type Pinger interface {
	Ping(ctx context.Context) error
}

// Query is a named query struct.
type Query struct {
	Name     string
	QueryRaw string
}

// TxManager is an interface for transaction manager, 
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

