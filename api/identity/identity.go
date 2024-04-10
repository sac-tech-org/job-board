package identity

import (
	"fmt"

	"github.com/rusher2004/job-board/api/server"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
)

type IdentityStore struct{}

func (i IdentityStore) GetUser(id string) (server.UserIdentity, error) {
	u, err := thirdpartyemailpassword.GetUserById(id)
	if err != nil {
		return server.UserIdentity{}, fmt.Errorf("error getting user: %w", err)
	}

	return server.UserIdentity{
		Email: server.Email{
			Address: u.Email,
		},
		ID: u.ID,
	}, nil
}

func NewIdentityStore() IdentityStore {
	return IdentityStore{}
}
