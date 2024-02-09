package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenConnSqlxPq(connString string) *sqlx.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db := sqlx.MustConnect("postgres", connString)

	err := db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return db
}

func OpenConnSqlx(connString string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("errors open connect mysql: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("errors ping connect mysql: %w", err)
	}

	return db, nil
}
