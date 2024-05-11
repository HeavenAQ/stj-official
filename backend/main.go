package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"stj-ecommerce/api"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func getDatabaseURL(config utils.Config) string {
	dbSource := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBDriver,
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSSLMode,
	)
	return dbSource
}

func setupDatabaseStore(config utils.Config) *db.Store {
	ctx := context.Background()
	dbSource := getDatabaseURL(config)
	testDBPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	store := db.NewStore(testDBPool)
	return store
}

func main() {
	// load environment variables
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	// set up database
	store := setupDatabaseStore(config)
	defer store.Close()

	// set up server
	server := api.NewServer(store, config)
	server.Start()
}
