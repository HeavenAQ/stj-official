package helpers

import (
	"net/http"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/token"

	"github.com/gin-gonic/gin"
)

func AuthAndGetUser(ctx *gin.Context, req any, store *db.Store, authKey string) (*db.User, error) {
	// get user
	userID := ctx.MustGet(authKey).(*token.Payload).QueryID
	user, err := store.GetUserByID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return nil, err
	}
	return &user, nil
}
