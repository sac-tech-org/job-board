package userstore

import (
	"encoding/json"
	"fmt"

	"github.com/rusher2004/job-board/api/server"
	"github.com/supertokens/supertokens-golang/recipe/usermetadata"
)

type UserStore struct{}

func NewUserStore() UserStore {
	return UserStore{}
}

type UserMetadata struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

func (u UserStore) GetMetadata(id string) (server.UserMetadata, error) {
	res, err := usermetadata.GetUserMetadata(id)
	if err != nil {
		return server.UserMetadata{}, fmt.Errorf("error getting user metadata: %w", err)
	}

	out, err := mapToStruct[UserMetadata](res)
	if err != nil {
		return server.UserMetadata{}, fmt.Errorf("error mapping to struct: %w", err)
	}

	return server.UserMetadata(out), nil
}

func (u UserStore) UpdateMetadata(id string, um server.UserMetadata) (server.UserMetadata, error) {
	updateMap := map[string]any{}
	inMap := map[string]string{
		"first_name": um.FirstName,
		"last_name":  um.LastName,
		"username":   um.Username,
	}

	// only update fields that are not empty. Otherwise, the user metadata will be overwritten with empty values
	for k, v := range inMap {
		if v != "" {
			updateMap[k] = v
		}
	}

	res, err := usermetadata.UpdateUserMetadata(id, updateMap)
	if err != nil {
		return server.UserMetadata{}, fmt.Errorf("error updating user metadata: %w", err)
	}

	out, err := mapToStruct[UserMetadata](res)
	if err != nil {
		return server.UserMetadata{}, fmt.Errorf("error mapping to struct: %w", err)
	}

	return server.UserMetadata(out), nil
}

func mapToStruct[T any](m map[string]any) (T, error) {
	var empty T

	b, err := json.Marshal(m)
	if err != nil {
		return empty, fmt.Errorf("error marshalling map: %w", err)
	}

	var out T
	if err := json.Unmarshal(b, &out); err != nil {
		return empty, fmt.Errorf("error unmarshalling map: %w", err)
	}

	return out, nil
}
