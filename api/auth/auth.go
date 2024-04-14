package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/recipe/usermetadata"
	"github.com/supertokens/supertokens-golang/supertokens"
)

type DBClient interface {
	InsertUser(ctx context.Context, first, id, last, username string) error
	UsernameExists(context.Context, string) (bool, error)
}

type AuthConfig struct {
	APIDomain  string
	APIKey     string
	AppName    string
	BasePath   string
	ConnURI    string
	SiteDomain string
	SiteURI    string
}

type SocialConfig struct {
	ProviderName string
	ClientID     string
	ClientSecret string
}

type AuthStore struct {
	cfg     AuthConfig
	db      DBClient
	socials []SocialConfig
	webURI  string
}

func (a *AuthStore) GetProviders() []tpmodels.ProviderInput {
	var p = make([]tpmodels.ProviderInput, len(a.socials))

	for i, s := range a.socials {
		p[i] = tpmodels.ProviderInput{
			Config: tpmodels.ProviderConfig{
				ThirdPartyId: s.ProviderName,
				Name:         s.ProviderName,
				Clients: []tpmodels.ProviderClientConfig{
					{
						ClientID:     s.ClientID,
						ClientSecret: s.ClientSecret,
					},
				},
			},
		}
	}

	return p
}

func (a *AuthStore) GetTypeInput() supertokens.TypeInput {
	return supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			APIKey:        a.cfg.APIKey,
			ConnectionURI: a.cfg.ConnURI,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         a.cfg.AppName,
			APIDomain:       a.cfg.APIDomain,
			WebsiteDomain:   a.cfg.SiteDomain,
			APIBasePath:     &a.cfg.BasePath,
			WebsiteBasePath: &a.cfg.BasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailverification.Init(evmodels.TypeInput{
				Mode: evmodels.ModeRequired,
			}),
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				Providers: a.GetProviders(),
				Override: &tpepmodels.OverrideStruct{
					APIs: a.apiOverrides,
				},
			}),
			session.Init(nil),
			dashboard.Init(nil),
			usermetadata.Init(nil),
		},
	}
}

func (a *AuthStore) apiOverrides(ogImplementation tpepmodels.APIInterface) tpepmodels.APIInterface {
	ogEPSignUp := *ogImplementation.EmailPasswordSignUpPOST

	(*ogImplementation.EmailPasswordSignUpPOST) = func(formFields []epmodels.TypeFormField, tenantId string, options epmodels.APIOptions, uc supertokens.UserContext) (tpepmodels.SignUpPOSTResponse, error) {
		// pull our custom user context from the request to do username uniqueness check and insert user
		type reqBody struct {
			UserContext struct {
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
				Username  string `json:"username"`
			} `json:"userContext"`
		}

		req := supertokens.GetRequestFromUserContext(uc)
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return tpepmodels.SignUpPOSTResponse{}, fmt.Errorf("error reading request body: %w", err)
		}

		ctx := req.Context()

		var u reqBody
		if err := json.Unmarshal(b, &u); err != nil {
			return tpepmodels.SignUpPOSTResponse{}, fmt.Errorf("error unmarshalling request body: %w", err)
		}

		exists, err := a.db.UsernameExists(ctx, u.UserContext.Username)
		if err != nil {
			return tpepmodels.SignUpPOSTResponse{}, fmt.Errorf("error checking if username exists: %w", err)
		}

		if exists {
			return tpepmodels.SignUpPOSTResponse{}, errors.New(`{"error": "username already exists"}`)
		}

		// here we just let supertokens to its normal thing.
		resp, err := ogEPSignUp(formFields, tenantId, options, uc)
		if err != nil {
			return tpepmodels.SignUpPOSTResponse{}, fmt.Errorf("error calling original email password sign up: %w", err)
		}

		if resp.OK != nil {
			if err := a.db.InsertUser(ctx, u.UserContext.FirstName, resp.OK.User.ID, u.UserContext.LastName, u.UserContext.Username); err != nil {
				return tpepmodels.SignUpPOSTResponse{}, fmt.Errorf("error inserting user: %w", err)
			}
		}

		return resp, err
	}

	return ogImplementation
}

func (a *AuthStore) GetWebURI() string {
	return a.webURI
}

func NewAuthStore(ac AuthConfig, db DBClient, sc []SocialConfig) AuthStore {
	return AuthStore{
		cfg:     ac,
		db:      db,
		socials: sc,
		webURI:  ac.SiteURI,
	}
}
