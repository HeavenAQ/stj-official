package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"tsn-ecommerce/api"
	db "tsn-ecommerce/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func getDatabaseURL() string {
	dbSource := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)
	return dbSource
}

func setupDatabaseStore() *db.Store {
	ctx := context.Background()
	dbSource := getDatabaseURL()
	testDBPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	store := db.NewStore(testDBPool)
	return store
}

func main() {
	// load .env file
	godotenv.Load()

	// set up database
	store := setupDatabaseStore()
	defer store.Close()

	// set up server
	server := api.NewServer(store)
	server.Start()
}
