package db

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rusher2004/job-board/api/datastore"
)

type User struct {
	CreatedAt time.Time
	ID        string
	FirstName string
	LastName  string
	UpdatedAt time.Time
	Username  string
	UUID      uuid.UUID
}

var ErrUserNotFound = errors.New("user not found")

func (c *Client) InsertUser(ctx context.Context, first, id, last, username string) error {
	userQuery := `
		INSERT INTO users.user (first_name, last_name, username, user_uuid)
		VALUES ($1, $2, $3, $4);
	`

	return pgx.BeginFunc(ctx, c.db, func(tx pgx.Tx) error {
		com, err := tx.Exec(ctx, userQuery, first, last, username, id)
		if err != nil {
			return fmt.Errorf("error inserting user: %w", err)
		}

		if com.RowsAffected() != 1 {
			return fmt.Errorf("expected 1 row affected, got %d", com.RowsAffected())
		}

		return nil
	})
}

func (c *Client) QueryUserByID(ctx context.Context, id string) (datastore.User, error) {
	query := `
		SELECT
			user_uuid,
			first_name,
			last_name,
			username
		FROM users.user
		WHERE user_uuid = $1
	`
	row := c.db.QueryRow(ctx, query, id)

	var user User
	if err := row.Scan(&user.UUID, &user.FirstName, &user.LastName, &user.Username); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return datastore.User{}, ErrUserNotFound
		}
		return datastore.User{}, fmt.Errorf("error scanning user row: %w", err)
	}

	return datastore.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		ID:        user.UUID.String(),
	}, nil
}

func (c *Client) UpdateUser(ctx context.Context, first, id, last, username string) error {
	valMap := map[string]string{
		"first_name": first,
		"last_name":  last,
		"username":   username,
	}

	var setLabels []string
	var setVals []any
	labelIDX := 1
	for k, v := range valMap {
		if v != "" {
			setLabels = append(setLabels, fmt.Sprintf("%s = $%d", k, labelIDX))
			setVals = append(setVals, v)
			labelIDX++
		}
	}

	if len(setLabels) == 0 {
		return errors.New("no values to update")
	}

	setClause := "SET " + strings.Join(setLabels, ", ") + "\n"
	whereClause := fmt.Sprintf("WHERE user_uuid = $%d;", labelIDX)
	query := "UPDATE users.user\n" + setClause + whereClause
	setVals = append(setVals, id)

	return pgx.BeginFunc(ctx, c.db, func(tx pgx.Tx) error {
		com, err := tx.Exec(ctx, query, setVals...)
		if err != nil {
			return fmt.Errorf("error updating user: %w", err)
		}

		if com.RowsAffected() != 1 {
			return fmt.Errorf("expected 1 row affected, got %d", com.RowsAffected())
		}

		return nil
	})
}

func (c *Client) UsernameExists(ctx context.Context, username string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users.user
			WHERE username = $1
		)
	`

	var exists bool
	if err := c.db.QueryRow(ctx, query, username).Scan(&exists); err != nil {
		return false, fmt.Errorf("error checking if username exists: %w", err)
	}

	return exists, nil
}
