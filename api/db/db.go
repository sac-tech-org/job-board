package db

import (
	"context"
	"fmt"
	"net/url"
	"strings"

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

	if !strings.Contains(u.String(), "/board") {
		u = u.JoinPath("board")
	}

	cfg, err := pgxpool.ParseConfig(u.String())
	if err != nil {
		return Client{}, fmt.Errorf("error parsing db config: %w", err)
	}

	p, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return Client{}, fmt.Errorf("error creating db pool: %w", err)
	}

	return Client{p}, nil
}

func (c *Client) Close() {
	c.db.Close()
}
