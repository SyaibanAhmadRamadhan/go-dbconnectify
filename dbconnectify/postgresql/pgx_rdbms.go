package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	godbconnectify "github.com/SyaibanAhmadRamadhan/go-dbconnectify"
)

type Commander interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

type pgxCommander struct {
	pool *pgxpool.Pool
}

func (r *pgxCommander) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.Exec(ctx, sql, arguments...)
	}
	return r.pool.Exec(ctx, sql, arguments...)
}

func (r *pgxCommander) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.Query(ctx, sql, args...)
	}
	return r.pool.Query(ctx, sql, args...)
}

func (r *pgxCommander) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.QueryRow(ctx, sql, args...)
	}
	return r.pool.QueryRow(ctx, sql, args...)
}

func (r *pgxCommander) Begin(ctx context.Context) (pgx.Tx, error) {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.Begin(ctx)
	}
	return r.Begin(ctx)
}

func (r *pgxCommander) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	tx := r.extractTx(ctx)
	if tx != nil {
		return tx.SendBatch(ctx, b)
	}
	return r.pool.SendBatch(ctx, b)
}

func (r *pgxCommander) extractTx(ctx context.Context) pgx.Tx {
	val := ctx.Value(godbconnectify.TxKey{})

	if tx, ok := val.(pgx.Tx); tx != nil && ok {
		return tx
	}

	return nil
}

// TX IMPLEMENTATION
type txPgx struct {
	pool *pgxpool.Pool
}

func NewTxPgx(pool *pgxpool.Pool) godbconnectify.Tx {
	return &txPgx{
		pool: pool,
	}
}

func (t *txPgx) DoTransaction(ctx context.Context, opt *godbconnectify.TxOption, fn func(c context.Context) (commit bool, err error)) (err error) {
	opts, err := t.extractOpt(opt)
	if err != nil {
		return err
	}

	var tx pgx.Tx
	if opts.IsoLevel == "" {
		tx, err = t.pool.Begin(ctx)
	} else {
		tx, err = t.pool.BeginTx(ctx, opts)
	}
	if err != nil {
		return err
	}

	var commit bool
	defer func() {
		if p := recover(); p != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Join(godbconnectify.ErrRollback, errRollback)
			}
			panic(p)
		} else if commit {
			if errCommit := tx.Commit(ctx); errCommit != nil {
				err = errors.Join(godbconnectify.ErrCommit, errCommit)
			}
		} else if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Join(godbconnectify.ErrRollback, errRollback)
			}
		} else {
			if errCommit := tx.Commit(ctx); errCommit != nil {
				err = errors.Join(godbconnectify.ErrCommit, errCommit)
			}
		}
	}()

	txKey := context.WithValue(ctx, godbconnectify.TxKey{}, tx)

	commit, err = fn(txKey)

	return err
}

func (t *txPgx) DoTransactionx(ctx context.Context, opt *godbconnectify.TxOption, fn func(c context.Context) (err error)) (err error) {
	opts, err := t.extractOpt(opt)
	if err != nil {
		return err
	}

	var tx pgx.Tx
	if opts.IsoLevel == "" {
		tx, err = t.pool.Begin(ctx)
	} else {
		tx, err = t.pool.BeginTx(ctx, opts)
	}
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Join(godbconnectify.ErrRollback, errRollback)
			}
			panic(p)
		} else if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Join(godbconnectify.ErrRollback, errRollback)
			}
		} else {
			if errCommit := tx.Commit(ctx); errCommit != nil {
				err = errors.Join(godbconnectify.ErrCommit, errCommit)
			}
		}
	}()

	txKey := context.WithValue(ctx, godbconnectify.TxKey{}, tx)

	err = fn(txKey)

	return err
}

func (t *txPgx) extractOpt(opt *godbconnectify.TxOption) (opts pgx.TxOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != godbconnectify.TxTypePgx && opt.Type != godbconnectify.TxTypeNone {
		err = fmt.Errorf("%w, your type is not *pgx.TxOptions. but %s", godbconnectify.ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(pgx.TxOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not pgx.TxOptions", godbconnectify.ErrTypeTx)
		return
	}

	return opts, nil
}
