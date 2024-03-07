package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rusher2004/job-board/api/datastore"
	"github.com/rusher2004/job-board/api/db"
	"github.com/rusher2004/job-board/api/server"

	_ "github.com/lib/pq"
)

func main() {
	dsn, err := getDBConfig()
	if err != nil {
		log.Fatalf("error getting dsn: %v", err)
	}
	db, err := db.NewDB(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	d := datastore.NewDataStore("some api client", &db)
	s := server.NewServer(d)

	log.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
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
