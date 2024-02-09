package mysql

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func MustOpenConnSqlx(connString string) *sqlx.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db := sqlx.MustConnect("mysql", connString)

	err := db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return db
}

func OpenConnSqlx(connString string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("errors open connect mysql: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("errors ping connect mysql: %w", err)
	}

	return db, nil
}
