package api

import (
	"errors"
	"fmt"
	"net/http"
	"stj-ecommerce/token"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func (server *Server) authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ensure the request has an authorization header
		authHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			server.ErrorLogger.Println(err)
			return
		}

		// ensure the field is in the correct format
		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := fmt.Errorf("invalid authorization header format: %s", authHeader)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			server.ErrorLogger.Println(err)
			return
		}

		// ensure the authorization type is bearer
		authType := strings.ToLower(fields[0])
		if authType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type: %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			server.ErrorLogger.Println(err)
			return
		}

		// ensure access token is valid
		payload, err := tokenMaker.VerifyToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			server.ErrorLogger.Println(err)
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
	}
}
