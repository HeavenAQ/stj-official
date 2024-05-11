package api

import (
	"net/http"
	db "tsn-ecommerce/db/sqlc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
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

func corsConfig() cors.Config {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	return corsConf
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	server.router = gin.Default()
	server.router.Use(cors.Default())
	server.router.Use(cors.New(corsConfig()))

	server.setupRouter()
	return server
}
