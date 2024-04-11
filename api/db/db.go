package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Client struct {
	db *pgxpool.Pool
}

func NewDB(ctx context.Context, connString string) (Client, error) {
	p, err := pgxpool.New(ctx, connString)
	if err != nil {
		return Client{}, fmt.Errorf("error creating db pool: %w", err)
	}

	return Client{p}, nil
}

func (c *Client) Close() {
	c.db.Close()
}
