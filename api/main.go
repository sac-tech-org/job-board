package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rusher2004/job-board/api/auth"
	"github.com/rusher2004/job-board/api/datastore"
	"github.com/rusher2004/job-board/api/db"
	"github.com/rusher2004/job-board/api/identity"
	"github.com/rusher2004/job-board/api/server"

	_ "github.com/lib/pq"
)

func main() {
	// setup auth store
	authCfg, err := getAuthConfig()
	if err != nil {
		log.Fatalf("error getting auth config: %v", err)
	}
	socCFG := getSocialConfigs()
	a := auth.NewAuthStore(authCfg, socCFG)

	// setup data store
	dsn, err := getDBConfig()
	if err != nil {
		log.Fatalf("error getting dsn: %v", err)
	}
	db, err := db.NewDB(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	d := datastore.NewDataStore(&db)

	// setup user store
	i := identity.NewIdentityStore()

	s, err := server.NewServer(&a, d, i)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}

func getAuthConfig() (auth.AuthConfig, error) {
	if err := godotenv.Load(); err != nil {
		return auth.AuthConfig{}, fmt.Errorf("error loading .env file: %w", err)
	}

	var (
		apiDomain  = os.Getenv("AUTH_API_DOMAIN")
		apiKey     = os.Getenv("AUTH_API_KEY")
		appName    = os.Getenv("AUTH_APP_NAME")
		basePath   = os.Getenv("AUTH_BASE_PATH")
		connURI    = os.Getenv("AUTH_CONN_URI")
		siteDomain = os.Getenv("AUTH_SITE_DOMAIN")
		siteURI    = os.Getenv("AUTH_SITE_URI")
	)

	if apiDomain == "" || apiKey == "" || appName == "" || basePath == "" || connURI == "" || siteDomain == "" || siteURI == "" {
		return auth.AuthConfig{}, fmt.Errorf("some auth config is missing")
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

func getDBConfig() (db.ConnectConfig, error) {
	if err := godotenv.Load(); err != nil {
		return db.ConnectConfig{}, fmt.Errorf("error loading .env file: %w", err)
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	if host == "" || port == "" || user == "" || password == "" {
		return db.ConnectConfig{}, fmt.Errorf("some db config is missing")
	}

	return db.ConnectConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}, nil
}

func getSocialConfigs() []auth.SocialConfig {
	return []auth.SocialConfig{
		{
			ProviderName: "github",
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		},
		{
			ProviderName: "google",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		},
		{
			ProviderName: "linkedin",
			ClientID:     os.Getenv("LINKEDIN_CLIENT_ID"),
			ClientSecret: os.Getenv("LINKEDIN_CLIENT_SECRET"),
		},
	}
}
