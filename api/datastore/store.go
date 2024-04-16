package datastore

import (
	"context"
)

type DBClient interface {
	InsertUser(ctx context.Context, first, id, last, username string) error
	QueryUserByID(context.Context, string) (User, error)
	UpdateUser(ctx context.Context, first, id, last, username string) error
}

type DataStore struct {
	db DBClient
}

func NewDataStore(db DBClient) DataStore {
	return DataStore{db}
}
