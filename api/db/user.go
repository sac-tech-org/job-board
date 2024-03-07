package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rusher2004/job-board/api/datastore"
)

type Email struct {
	Address   string
	CreatedAt time.Time
	ID        int
	UpdatedAt time.Time
	UUID      uuid.UUID
	Verified  bool
}

type User struct {
	CreatedAt time.Time
	Email     Email
	ID        string
	FirstName string
	LastName  string
	UpdatedAt time.Time
	Username  string
	UUID      uuid.UUID
}

var ErrUserNotFound = errors.New("user not found")

func (c *Client) InsertUser(ctx context.Context, email, first, last, username string) (datastore.User, error) {
	userQuery := `
		INSERT INTO users.user (first_name, last_name, username)
		VALUES ($1, $2, $3)
		RETURNING id, user_uuid;
	`

	emailQuery := `
		INSERT INTO users.email (address, user_id, primary)
		VALUES ($1, $2, true);
	`

	var outUUID string
	if err := pgx.BeginFunc(ctx, c.db, func(tx pgx.Tx) error {
		var id string
		if err := tx.QueryRow(ctx, userQuery, first, last, username).Scan(&id, &outUUID); err != nil {
			return fmt.Errorf("error inserting user: %w", err)
		}

		com, err := tx.Exec(ctx, emailQuery, email, id)
		if err != nil {
			return fmt.Errorf("error inserting email: %w", err)
		}

		if com.RowsAffected() != 1 {
			return fmt.Errorf("expected 1 row affected, got %d", com.RowsAffected())
		}

		return nil
	}); err != nil {
		return datastore.User{}, fmt.Errorf("error executing transaction: %w", err)
	}

	return datastore.User{
		Email:     email,
		FirstName: first,
		LastName:  last,
		Username:  username,
		UUID:      uuid.MustParse(outUUID),
	}, nil
}

func (c *Client) QueryUserByID(ctx context.Context, id uuid.UUID) (datastore.User, error) {
	query := `
		SELECT
			u.user_uuid,
			u.first_name,
			u.last_name,
			u.username,
			e.address
		FROM users.user u
		INNER JOIN users.email e ON e.user_id = u.id
		WHERE u.user_uuid = $1
	`
	row := c.db.QueryRow(ctx, query, id)

	var user User
	if err := row.Scan(&user.UUID, &user.FirstName, &user.LastName, &user.Username, &user.Email.Address); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return datastore.User{}, ErrUserNotFound
		}
		return datastore.User{}, fmt.Errorf("error scanning user row: %w", err)
	}

	return datastore.User{
		Email:     user.Email.Address,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		UUID:      user.UUID,
	}, nil
}
