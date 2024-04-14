package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbURL, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		log.Fatal("POSTGRES_URL is not set")
	}

	ghFile, ok := os.LookupEnv("GITHUB_STEP_SUMMARY")
	if !ok {
		log.Fatal("GITHUB_STEP_SUMMARY is not set")
	}

	fmt.Printf("dbURL: %v\n", dbURL)
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatalf("error creating migrator: %v\n", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("error applying migrations: %v\n", err)
	}

	os.WriteFile(ghFile, []byte("migrations applied"), 0644)
}
