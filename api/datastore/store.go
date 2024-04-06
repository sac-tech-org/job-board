package datastore

import (
	"context"

	"github.com/google/uuid"
)

type DBClient interface {
	InsertUser(ctx context.Context, email, first, last, username string) (User, error)
	QueryUserByID(context.Context, uuid.UUID) (User, error)
}

type DataStore struct {
	db DBClient
}

func NewDataStore(db DBClient) DataStore {
	return DataStore{db}
}
