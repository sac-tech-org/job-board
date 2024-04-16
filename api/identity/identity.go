package identity

import (
	"fmt"

	"github.com/rusher2004/job-board/api/server"
	ev "github.com/supertokens/supertokens-golang/recipe/emailverification"
	tpep "github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
)

type IdentityStore struct{}

func (i IdentityStore) GetUser(id string) (server.UserIdentity, error) {
	u, err := tpep.GetUserById(id)
	if err != nil {
		return server.UserIdentity{}, fmt.Errorf("error getting user: %w", err)
	}

	v, err := ev.IsEmailVerified(id, &u.Email)
	if err != nil {
		return server.UserIdentity{}, fmt.Errorf("error getting email verification status: %w", err)
	}

	return server.UserIdentity{
		Email: server.Email{
			Address:  u.Email,
			Verified: v,
		},
		ID: u.ID,
	}, nil
}

func NewIdentityStore() IdentityStore {
	return IdentityStore{}
}
