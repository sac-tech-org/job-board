package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	db *pgxpool.Pool
}

type ConnectConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

func NewDB(ctx context.Context, cfg ConnectConfig) (Client, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password)
	p, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return Client{}, fmt.Errorf("error creating db pool: %w", err)
	}

	return Client{p}, nil
}

func (c *Client) Close() {
	c.db.Close()
}
