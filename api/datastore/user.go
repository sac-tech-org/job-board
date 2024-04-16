package datastore

import (
	"context"

	"github.com/rusher2004/job-board/api/server"
)

type User struct {
	FirstName string
	LastName  string
	ID        string
	Username  string
}

func (d DataStore) CreateUser(ctx context.Context, in server.PostUserInput) error {
	return d.db.InsertUser(ctx, in.FirstName, in.ID, in.LastName, in.Username)
}

func (d DataStore) DeleteUser(context.Context, server.DeleteUserInput) error {
	return nil
}

func (d DataStore) GetUser(ctx context.Context, in server.GetUserInput) (server.UserMetadata, error) {
	u, err := d.db.QueryUserByID(ctx, in.ID)
	if err != nil {
		return server.UserMetadata{}, err
	}

	return server.UserMetadata{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
	}, nil
}

func (d DataStore) GetUserList(ctx context.Context, in server.GetUserListInput) ([]server.UserMetadata, error) {
	return []server.UserMetadata{}, nil
}

func (d DataStore) PutUser(ctx context.Context, first, id, last, username string) error {
	return d.db.UpdateUser(ctx, first, id, last, username)
}
