package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/token"
	"stj-ecommerce/utils"

	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	config      utils.Config
	store       *db.Store
	router      *gin.Engine
	tokenMaker  token.Maker
	ErrorLogger *log.Logger
	InforLogger *log.Logger
	WarnLogger  *log.Logger
}

func (server *Server) Start() error {
	return server.router.Run()
}

func (server *Server) healthCheck(ctx *gin.Context) {
	err := server.store.HealthCheck()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func setupServer(store *db.Store, config utils.Config) (*Server, error) {
	// setup token maker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSyemmetricKey)
	if err != nil {
		log.Println("cannot create token maker")
		log.Println(err)
		return nil, err
	}

	// config server
	server := &Server{store: store, config: config, tokenMaker: tokenMaker}
	server.router = gin.Default()
	server.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	server.setupRouter()
	return server, nil
}

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

func (server *Server) setupLogger() {
	red := color.New(color.FgRed)
	cyan := color.New(color.FgCyan)
	yellow := color.New(color.FgYellow)

	server.ErrorLogger = log.New(os.Stderr, red.Sprint("[ERROR] "), log.Ldate|log.Ltime|log.Llongfile)
	server.InforLogger = log.New(os.Stdout, cyan.Sprint("[INFO] "), log.Ldate|log.Ltime|log.Llongfile)
	server.WarnLogger = log.New(os.Stdout, yellow.Sprint("[WARN] "), log.Ldate|log.Ltime|log.Llongfile)
}

func NewServer(configPath string) (*Server, error) {
	// load .env file
	config, err := utils.LoadConfig(configPath)
	if err != nil {
		log.Printf(".env file not found: %v\n", err)
		log.Printf("Using default environment variables\n")
		config = utils.LoadConfigFromEnv()
	}

	// set up database
	store := setupDatabaseStore(config)

	// set up server
	server, err := setupServer(store, config)
	if err != nil {
		log.Println("cannot create server")
		return nil, err
	}

	// set up logger
	server.setupLogger()

	return server, err
}

func (server *Server) Shutdown() {
	server.store.Close()
}
