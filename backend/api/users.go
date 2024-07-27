package api

import (
	"net/http"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/helpers"
	apierrors "stj-ecommerce/helpers/api_errors"
	"stj-ecommerce/token"
	"stj-ecommerce/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserRequest struct {
	Email     string          `json:"email" binding:"required"`
	Gender    db.Gender       `json:"gender"`
	Password  string          `json:"password" binding:"required,min=8,max=100"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Language  db.LanguageCode `json:"language" binding:"required"`
	Address   string          `json:"address"`
	LineID    pgtype.Text     `json:"line_id"`
	Phone     pgtype.Text     `json:"phone"`
	Latitude  pgtype.Float8   `json:"latitude"`
	Longitude pgtype.Float8   `json:"longitude"`
	BirthYear pgtype.Int4     `json:"birth_year"`
}

type UserInfoResponse struct {
	Gender    db.Gender       `json:"gender"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Language  db.LanguageCode `json:"language"`
	Address   string          `json:"address"`
	LineID    pgtype.Text     `json:"line_id"`
	Latitude  pgtype.Float8   `json:"latitude"`
	Longitude pgtype.Float8   `json:"longitude"`
	BirthYear pgtype.Int4     `json:"birth_year"`
}

// Convert DB user to response to exclude sensitive information
func (server *Server) userResponse(user db.User) UserInfoResponse {
	return UserInfoResponse{
		Email:     user.Email,
		Phone:     user.Phone.String,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Language:  user.Language,
		Address:   user.Address,
		LineID:    user.LineID,
		BirthYear: user.BirthYear,
		Gender:    user.Gender,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
	}
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
		return
	}

	// decode password and hash it
	password, err := utils.DecodeAndHashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
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
		LineID:    req.LineID,
		Gender:    req.Gender,
		BirthYear: req.BirthYear,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	// create user in database
	user, err := server.store.CreateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusConflict, apierrors.UserRegistrationError(err))
		server.ErrorLogger.Println(err)
		return
	}

	// return user
	ctx.JSON(http.StatusOK, server.userResponse(user))
}

func (server *Server) GetUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUserByID(ctx, authPayload.QueryID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		server.ErrorLogger.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, server.userResponse(user))
}

type UpdateUserRequest struct {
	Email     string          `json:"email" binding:"required"`
	Gender    db.Gender       `json:"gender"`
	Password  string          `json:"password"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Language  db.LanguageCode `json:"language" binding:"required"`
	Address   string          `json:"address"`
	LineID    pgtype.Text     `json:"line_id"`
	Phone     pgtype.Text     `json:"phone"`
	Latitude  pgtype.Float8   `json:"latitude"`
	Longitude pgtype.Float8   `json:"longitude"`
	BirthYear pgtype.Int4     `json:"birth_year"`
}

func (server *Server) UpdateUser(ctx *gin.Context) {
	// ensure the request body is a valid JSON
	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get old user information
	oldUser, err := helpers.AuthAndGetUser(ctx, req, server.store, authorizationPayloadKey)
	if err != nil {
		server.ErrorLogger.Println(err)
	}

	// update user in database
	arg := db.UpdateUserByIdParams{
		ID:        oldUser.ID,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  oldUser.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Language:  req.Language,
		Address:   req.Address,
		LineID:    req.LineID,
		Gender:    req.Gender,
		BirthYear: req.BirthYear,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	user, err := server.store.UpdateUserById(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusConflict, apierrors.UserRegistrationError(err))
		server.ErrorLogger.Println(err)
		return
	}

	// return updated user
	ctx.JSON(http.StatusOK, server.userResponse(user))
}
