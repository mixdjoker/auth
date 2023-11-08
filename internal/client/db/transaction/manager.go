package transaction

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/mixdjoker/auth/internal/client/db"
	"github.com/mixdjoker/auth/internal/client/db/pg"
	"github.com/pkg/errors"
)

const (
	txBegErrMsg = "can't begin transaction"
	panicErrMsg = "panic recovered"
	commitErrMsg = "tx commit failed"
	rollbackErrMsg = "tx rollback failed"
	txErrMsg = "failed executing code inside transaction"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager returns a new instance of db.TxManager
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, txBegErrMsg)
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("%s: %v", panicErrMsg, r)
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "%s: %v", rollbackErrMsg, errRollback)
			}

			return
		}

		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, commitErrMsg)
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, txErrMsg)
	}

	return err
}
func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
