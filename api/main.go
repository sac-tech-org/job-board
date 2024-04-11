package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rusher2004/job-board/api/auth"
	"github.com/rusher2004/job-board/api/datastore"
	"github.com/rusher2004/job-board/api/db"
	"github.com/rusher2004/job-board/api/identity"
	"github.com/rusher2004/job-board/api/server"
)

func main() {
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := db.NewDB(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// setup auth store
	authCfg, err := getAuthConfig()
	if err != nil {
		log.Fatal(err)
	}
	socCFG, err := getSocialConfigs()
	if err != nil {
		log.Fatal(err)
	}
	a := auth.NewAuthStore(authCfg, &db, socCFG)

	// setup data store
	d := datastore.NewDataStore(&db)

	// setup user store
	i := identity.NewIdentityStore()

	s, err := server.NewServer(&a, d, i)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("----------------------------------")
	log.Println("            it changed            ")
	log.Println("----------------------------------")

	log.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}

func getAuthConfig() (auth.AuthConfig, error) {
	var envErrs []string
	apiDomain, ok := os.LookupEnv("AUTH_API_DOMAIN")
	if !ok {
		envErrs = append(envErrs, "AUTH_API_DOMAIN")
	}
	apiKey, ok := os.LookupEnv("AUTH_API_KEY")
	if !ok {
		envErrs = append(envErrs, "AUTH_API_KEY")
	}
	appName, ok := os.LookupEnv("AUTH_APP_NAME")
	if !ok {
		envErrs = append(envErrs, "AUTH_APP_NAME")
	}
	basePath, ok := os.LookupEnv("AUTH_BASE_PATH")
	if !ok {
		envErrs = append(envErrs, "AUTH_BASE_PATH")
	}
	connURI, ok := os.LookupEnv("AUTH_CONN_URI")
	if !ok {
		envErrs = append(envErrs, "AUTH_CONN_URI")
	}
	siteDomain, ok := os.LookupEnv("AUTH_SITE_DOMAIN")
	if !ok {
		envErrs = append(envErrs, "AUTH_SITE_DOMAIN")
	}
	siteURI, ok := os.LookupEnv("AUTH_SITE_URI")
	if !ok {
		envErrs = append(envErrs, "AUTH_SITE_URI")
	}

	if len(envErrs) > 0 {
		msg := "auth env vars missing: " + strings.Join(envErrs, "; ")
		return auth.AuthConfig{}, errors.New(msg)
	}
	return auth.AuthConfig{
		APIDomain:  apiDomain,
		APIKey:     apiKey,
		AppName:    appName,
		BasePath:   basePath,
		ConnURI:    connURI,
		SiteDomain: siteDomain,
		SiteURI:    siteURI,
	}, nil
}

func getSocialConfigs() ([]auth.SocialConfig, error) {
	var envErrs []string
	providers := []auth.SocialConfig{
		{ProviderName: "github"},
		{ProviderName: "google"},
		{ProviderName: "linkedin"},
	}

	for _, p := range providers {
		idVar := strings.ToUpper(p.ProviderName) + "_CLIENT_ID"
		secretVar := strings.ToUpper(p.ProviderName) + "_CLIENT_SECRET"
		clientID, ok := os.LookupEnv(idVar)
		if !ok {
			envErrs = append(envErrs, secretVar)
		}
		clientSecret, ok := os.LookupEnv(secretVar)
		if !ok {
			envErrs = append(envErrs, secretVar)
		}

		p.ClientID = clientID
		p.ClientSecret = clientSecret
	}

	if len(envErrs) > 0 {
		msg := "social env vars missing: " + strings.Join(envErrs, "; ")
		return nil, errors.New(msg)
	}

	return providers, nil
}
