package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenPgxPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)

	return conn, err
}

func OpenPgxPoolWithConfig(ctx context.Context, config *pgxpool.Config) (*pgxpool.Pool, error) {

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)

	return conn, err
}

func OpenPgxConn(ctx context.Context, connString string) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)

	return conn, err
}
