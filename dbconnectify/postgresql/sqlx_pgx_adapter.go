package postgresql

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func OpenConnSqlxPgx(connStr string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("errors open connect mysql: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("errors ping connect mysql: %w", err)
	}

	return db, nil
}
