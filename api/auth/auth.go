package auth

import (
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

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
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				Providers: a.GetProviders(),
			}),
			session.Init(nil),
			dashboard.Init(nil),
		},
	}
}

func (a *AuthStore) GetWebURI() string {
	return a.webURI
}

func NewAuthStore(ac AuthConfig, sc []SocialConfig) AuthStore {
	return AuthStore{
		cfg:     ac,
		socials: sc,
		webURI:  ac.SiteURI,
	}
}
