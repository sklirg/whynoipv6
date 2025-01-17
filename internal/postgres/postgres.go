package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPostgreSQL instantiates a new Postgres struct.
func NewPostgreSQL(conf string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, conf)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
