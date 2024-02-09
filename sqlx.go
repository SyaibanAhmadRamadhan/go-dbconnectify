package go_dbconnectify

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Sqlx struct {
	Commander Commander
	DB        *sqlx.DB
	Builder   squirrel.StatementBuilderType
}

func (s *Sqlx) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}

func NewSqlx(db *sqlx.DB) *Sqlx {
	return &Sqlx{
		Commander: &sqlxCommander{db: db},
		DB:        db,
		Builder:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
