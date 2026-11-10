package postgres

import (
	"context"
	"fmt"

	"github.com/asfqx/SubTrack/config"
	"github.com/asfqx/SubTrack/pkg/pgPool"
)

type Postgres struct {
	pool *pgPool.Pool
}

func New(ctx context.Context, cfg config.Config) (*Postgres, error) {
	pgPool, err := pgPool.New(ctx, cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("pgPool.New: %w", err)
	}

	if err := pgPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgPool ping: %w", err)
	}

	return &Postgres{
		pool: pgPool,
	}, nil
}

func (p *Postgres) Pool() *pgPool.Pool {
	return p.pool
}

func (p *Postgres) Close() {
	p.pool.Close()
}
