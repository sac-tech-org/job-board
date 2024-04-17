package datastore

import (
	"context"
)

func (d DataStore) CreateOrganization(ctx context.Context, descr, name, userID, website string) (string, error) {
	return d.db.InsertOrganization(ctx, descr, name, userID, website)
}
