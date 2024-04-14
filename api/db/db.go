package db

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	db *pgxpool.Pool
}

func NewDB(ctx context.Context, connString string) (Client, error) {
	u, err := url.Parse(connString)
	if err != nil {
		return Client{}, fmt.Errorf("error parsing db url: %w", err)
	}

	// explicitly set this because fly.io create another default db
	u.Path = "/board"
	cfg, err := pgxpool.ParseConfig(u.String())
	if err != nil {
		return Client{}, fmt.Errorf("error parsing db config: %w", err)
	}

	p, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return Client{}, fmt.Errorf("error creating db pool: %w", err)
	}

	if err := p.Ping(ctx); err != nil {
		return Client{}, fmt.Errorf("error pinging db: %w", err)
	}

	return Client{p}, nil
}

func (c *Client) Close() {
	c.db.Close()
}
