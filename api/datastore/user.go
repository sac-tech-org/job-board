package datastore

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rusher2004/job-board/api/server"
)

type User struct {
	Email     string
	FirstName string
	LastName  string
	UUID      uuid.UUID
	Username  string
}

func (d DataStore) DeleteUser(context.Context, server.DeleteUserInput) error {
	return nil
}

func (d DataStore) GetUser(ctx context.Context, in server.GetUserInput) (server.User, error) {
	u, err := d.db.QueryUserByID(ctx, in.UUID)
	if err != nil {
		return server.User{}, fmt.Errorf("error querying user by id: %w", err)
	}

	return server.User(u), nil
}

func (d DataStore) GetUserList(ctx context.Context, in server.GetUserListInput) ([]server.User, error) {
	return []server.User{}, nil
}

func (d DataStore) PostUser(ctx context.Context, in server.PostUserInput) (server.User, error) {
	u, err := d.db.InsertUser(ctx, in.Email, in.FirstName, in.LastName, in.Username)
	if err != nil {
		return server.User{}, fmt.Errorf("error creating user: %w", err)
	}

	return server.User(u), nil
}

func (d DataStore) PutUser(ctx context.Context, in server.PutUserInput) (server.User, error) {
	return server.User{}, nil
}
