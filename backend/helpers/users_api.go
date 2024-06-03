package helpers

import (
	"net/http"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyJSONAndGetUser(ctx *gin.Context, req any, store *db.Store, authKey string) (*db.User, error) {
	// ensure the request body is a valid JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	// get user
	userID := ctx.MustGet(authKey).(*token.Payload).QueryID
	user, err := store.GetUserByID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return nil, err
	}
	return &user, nil
}

func UserErrorResponse(err error) gin.H {
	if strings.Contains(err.Error(), "users_email_key") {
		return gin.H{"error": "email already exists"}
	} else if strings.Contains(err.Error(), "users_phone_key") {
		return gin.H{"error": "phone already exists"}
	} else if strings.Contains(err.Error(), "users_line_id_key") {
		return gin.H{"error": "line id already exists"}
	}
	return gin.H{"error": err.Error()}
}
