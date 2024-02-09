package go_dbconnectify

import (
	"context"
	"errors"
)

type TxType uint

const (
	TxTypeNone TxType = iota
	TxTypeMongoDB
	TxTypePgx
	TxTypeSqlx
)

var ErrRollback = errors.New("error rollback db | ")
var ErrCommit = errors.New("error commit db | ")

func (t TxType) String() string {
	switch t {
	case TxTypeNone:
		return "none type"
	case TxTypeMongoDB:
		return "mongodb"
	case TxTypePgx:
		return "pgx"
	case TxTypeSqlx:
		return "sqlx"
	}

	return ""
}

var ErrTypeTx = errors.New("invalid type tx")

type TxOption struct {
	Type   TxType
	Option interface{}
}

type TxKey struct{}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . Tx
type Tx interface {
	DoTransaction(ctx context.Context, opt *TxOption, fn func(c context.Context) (commit bool, err error)) (err error)
	DoTransactionx(ctx context.Context, opt *TxOption, fn func(c context.Context) (err error)) (err error)
}
