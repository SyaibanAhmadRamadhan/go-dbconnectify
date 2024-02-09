package postgresql

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPgx struct {
	Commander Commander
	Builder   squirrel.StatementBuilderType
	Pool      *pgxpool.Pool
}

func (p *PostgresPgx) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func NewPgxPostgres(pool *pgxpool.Pool) *PostgresPgx {
	return &PostgresPgx{
		Commander: &pgxCommander{pool: pool},
		Builder:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		Pool:      pool,
	}
}
