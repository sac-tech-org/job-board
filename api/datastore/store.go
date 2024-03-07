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
	lambdaClient any
	db           DBClient
}

func NewDataStore(cl any, db DBClient) DataStore {
	return DataStore{cl, db}
}
