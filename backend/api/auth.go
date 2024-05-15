package api

import (
	"database/sql"
	"net/http"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	SessionID           uuid.UUID        `json:"session_id"`
	AccessToken         string           `json:"access_token"`
	RefreshToken        string           `json:"refresh_token"`
	AccessTokenExpires  int64            `json:"access_token_expires"`
	RefreshTokenExpires int64            `json:"refresh_token_expires"`
	User                UserInfoResponse `json:"user"`
}

func (server *Server) UserLogin(ctx *gin.Context) {
	var req loginRequest
	// ensure the request is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
		return
	}

	// check if the user exists
	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			server.ErrorLogger.Println(err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
		return
	}

	// check if the password is correct
	err = utils.CheckPassword(user.Password, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
		return
	}

	// create access token
	accessToken, _, err := server.tokenMaker.CreateToken(
		user.ID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
		return
	}

	// create refresh token
	refreshToken, payload, err := server.tokenMaker.CreateToken(
		user.ID,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
		return
	}

	// store session in database (NOTE: could migrate to redis in the future)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt: pgtype.Timestamptz{
			Time:  payload.ExpiresAt.Time,
			Valid: true,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
		return
	}

	// create response
	ctx.JSON(http.StatusOK, loginResponse{
		SessionID:           session.ID.Bytes,
		AccessToken:         accessToken,
		RefreshToken:        session.RefreshToken,
		AccessTokenExpires:  payload.ExpiresAt.Time.Unix(),
		RefreshTokenExpires: payload.ExpiresAt.Time.Unix(),
		User:                server.userResponse(user),
	})
}

func (server *Server) UserRegister(ctx *gin.Context) {
	// TODO: email verification
	server.CreateUser(ctx)
}
