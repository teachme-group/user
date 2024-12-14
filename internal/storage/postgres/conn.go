package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Conn interface {
		Queries(ctx context.Context) *Queries
		//WithTx(ctx context.Context, txFunc func(ctx context.Context) error) error
	}

	conn struct {
		conn    *pgxpool.Pool
		queries *Queries
	}
)

func NewStorage(
	pool *pgxpool.Pool,
) *conn {
	return &conn{
		conn:    pool,
		queries: New(pool),
	}
}

func (t *conn) Queries(ctx context.Context) *Queries {
	tx := extractTx(ctx)
	if tx != nil {
		return t.queries.WithTx(tx)
	}

	return t.queries
}
