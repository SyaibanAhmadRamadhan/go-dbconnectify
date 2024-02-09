package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	godbconnectify "github.com/SyaibanAhmadRamadhan/go-dbconnectify"
)

type txMongodb struct {
	client *mongo.Client
}

func NewTxMongodb(client *mongo.Client) godbconnectify.Tx {
	return &txMongodb{client: client}
}

func (m *txMongodb) DoTransaction(ctx context.Context, opt *godbconnectify.TxOption, fn func(c context.Context) (commit bool, err error)) (err error) {
	opts, err := m.extractOpt(opt)
	if err != nil {
		return
	}

	var session mongo.Session
	if opts == nil {
		session, err = m.client.StartSession()
	} else {
		session, err = m.client.StartSession(opts)
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = session.StartTransaction()
		if err != nil {
			return err
		}

		var commit bool
		defer func() {
			if p := recover(); p != nil {
				if errRollback := session.AbortTransaction(sc); errRollback != nil {
					err = errors.Join(godbconnectify.ErrRollback, errRollback)
				}
				panic(p)
			} else if commit {
				if errCommit := session.CommitTransaction(sc); errCommit != nil {
					err = errors.Join(godbconnectify.ErrCommit, errCommit)
				}
			} else if err != nil {
				if errRollback := session.AbortTransaction(sc); errRollback != nil {
					err = errors.Join(godbconnectify.ErrRollback, errRollback)
				}
			} else {
				if errCommit := session.CommitTransaction(sc); errCommit != nil {
					err = errors.Join(godbconnectify.ErrCommit, errCommit)
				}
			}
		}()

		commit, err = fn(sc)
		return err
	})

	return err
}
func (m *txMongodb) DoTransactionx(ctx context.Context, opt *godbconnectify.TxOption, fn func(c context.Context) (err error)) (err error) {
	opts, err := m.extractOpt(opt)
	if err != nil {
		return
	}

	var session mongo.Session
	if opts == nil {
		session, err = m.client.StartSession()
	} else {
		session, err = m.client.StartSession(opts)
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = session.StartTransaction()
		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				if errRollback := session.AbortTransaction(sc); errRollback != nil {
					err = errors.Join(godbconnectify.ErrRollback, errRollback)
				}
				panic(p)
			} else if err != nil {
				if errRollback := session.AbortTransaction(sc); errRollback != nil {
					err = errors.Join(godbconnectify.ErrRollback, errRollback)
				}
			} else {
				if errCommit := session.CommitTransaction(sc); errCommit != nil {
					err = errors.Join(godbconnectify.ErrCommit, errCommit)
				}
			}
		}()

		err = fn(sc)
		return err
	})

	return err
}

func (m *txMongodb) extractOpt(opt *godbconnectify.TxOption) (opts *options.SessionOptions, err error) {
	if opt == nil {
		return
	}

	if opt.Option == nil {
		return
	}

	if opt.Type != godbconnectify.TxTypeMongoDB && opt.Type != godbconnectify.TxTypeNone {
		err = fmt.Errorf("%w, your type is not pgx. but %s", godbconnectify.ErrTypeTx, opt.Type.String())
		return
	}

	opts, ok := opt.Option.(*options.SessionOptions)
	if !ok {
		err = fmt.Errorf("%w, your type is not *options.SessionOptions", godbconnectify.ErrTypeTx)
		return
	}

	return opts, nil
}
