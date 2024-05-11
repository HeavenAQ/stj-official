package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var testQueries *Queries
var testStore *Store

func TestMain(m *testing.M) {
	// load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf(".env file not found: %v\n", err)
		log.Printf("Using default environment variables\n")
	}

	// set up connection pool
	ctx := context.Background()
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
	testDBPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer testDBPool.Close()

	// set up queries and run
	testQueries = New(testDBPool)
	testStore = NewStore(testDBPool)
	m.Run()
}
