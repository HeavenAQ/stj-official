package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"stj-ecommerce/utils"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testStore *Store

func TestMain(m *testing.M) {
	// load .env file
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Printf(".env file not found: %v\n", err)
		log.Printf("Using default environment variables\n")
		config = utils.LoadConfigFromEnv()
	}

	// set up connection pool
	ctx := context.Background()
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
