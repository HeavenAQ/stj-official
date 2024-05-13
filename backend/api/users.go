package api

import (
	"fmt"
	"net/http"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/token"
	"stj-ecommerce/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserRequest struct {
	Email     string          `json:"email" binding:"required"`
	Phone     pgtype.Text     `json:"phone"`
	Password  string          `json:"password" binding:"required,min=8,max=100"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Language  db.LanguageCode `json:"language" binding:"required"`
	Address   string          `json:"address"`
}

type UserInfoResponse struct {
	ID        pgtype.UUID     `json:"id"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Language  db.LanguageCode `json:"language"`
	Address   string          `json:"address"`
}

// Convert DB user to response to exclude sensitive information
func (server *Server) userResponse(user db.User) UserInfoResponse {
	return UserInfoResponse{
		ID:        user.ID,
		Email:     user.Email,
		Phone:     user.Phone.String,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Language:  user.Language,
		Address:   user.Address,
	}
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// decode password and hash it
	password, err := utils.DecodeAndHashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create args for creating user
	args := db.CreateUserParams{
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Language:  req.Language,
		Address:   req.Address,
	}

	// create user in database
	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			ctx.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// return user
	ctx.JSON(http.StatusOK, server.userResponse(user))
}

func (server *Server) GetUser(ctx *gin.Context) {
	fmt.Println("GetUser")
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUserById(ctx, authPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, server.userResponse(user))
}
