package go_dbconnectify

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Commander interface {
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
}

type sqlxCommander struct {
	db *sqlx.DB
}

func (c *sqlxCommander) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.QueryxContext(ctx, query, args...)
	}

	return c.db.QueryxContext(ctx, query, args...)
}

func (c *sqlxCommander) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.QueryRowxContext(ctx, query, args...)
	}

	return c.db.QueryRowxContext(ctx, query, args...)
}

func (c *sqlxCommander) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}

	return c.db.ExecContext(ctx, query, args...)
}

func (c *sqlxCommander) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.NamedQuery(query, arg)
	}

	return c.db.NamedQueryContext(ctx, query, arg)
}

func (c *sqlxCommander) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.NamedExecContext(ctx, query, arg)
	}

	return c.db.NamedExecContext(ctx, query, arg)
}

func (c *sqlxCommander) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.PreparexContext(ctx, query)
	}

	return c.db.PreparexContext(ctx, query)
}

func (c *sqlxCommander) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	tx := c.extractTx(ctx)
	if tx != nil {
		return tx.PrepareNamedContext(ctx, query)
	}

	return c.db.PrepareNamedContext(ctx, query)
}

func (c *sqlxCommander) extractTx(ctx context.Context) (tx *sqlx.Tx) {
	val := ctx.Value(TxKey{})

	if tx, ok := val.(*sqlx.Tx); tx != nil && ok {
		return tx
	}

	return nil
}

// TX SQLX IMPLEMENTATION
type txSqlx struct {
	db *sqlx.DB
}

func NewTxSqlx(db *sqlx.DB) Tx {
	return &txSqlx{
		db: db,
	}
}

func (t *txSqlx) DoTransaction(ctx context.Context, opt *TxOption, fn func(c context.Context) (commit bool, err error)) (err error) {
	opts, err := t.extractOpt(opt)
	if err != nil {
		return err
	}

	var tx *sqlx.Tx
	if opts == nil {
		tx, err = t.db.Beginx()
	} else {
		tx, err = t.db.BeginTxx(ctx, opts)
	}
	if err != nil {
		return err
	}

	var commit bool

	defer func() {
		if p := recover(); p != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = errors.Join(ErrRollback, errRollback)
			}
			panic(p)
		} else if commit {
			if errCommit := tx.Commit(); errCommit != nil {
				err = errors.Join(ErrCommit, errCommit)
			}
		} else if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = errors.Join(ErrRollback, errRollback)
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				err = errors.Join(ErrCommit, errCommit)
			}
		}
	}()

	txKey := context.WithValue(ctx, TxKey{}, tx)

	commit, err = fn(txKey)

	return err
}

func (t *txSqlx) DoTransactionx(ctx context.Context, opt *TxOption, fn func(c context.Context) (err error)) (err error) {
	opts, err := t.extractOpt(opt)
	if err != nil {
		return err
	}

	var tx *sqlx.Tx
	if opts == nil {
		tx, err = t.db.Beginx()
	} else {
		tx, err = t.db.BeginTxx(ctx, opts)
	}
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = errors.Join(ErrRollback, errRollback)
			}
			panic(p)
		} else if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = errors.Join(ErrRollback, errRollback)
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				err = errors.Join(ErrCommit, errCommit)
			}
		}
	}()

	txKey := context.WithValue(ctx, TxKey{}, tx)

	err = fn(txKey)

	return err
}

func (t *txSqlx) extractOpt(opt *TxOption) (opts *sql.TxOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != TxTypeSqlx && opt.Type != TxTypeNone {
		err = fmt.Errorf("%w, your type is not *sql.TxOptions. but %s", ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(*sql.TxOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not *sql.TxOptions", ErrTypeTx)
		return
	}

	return opts, nil
}
