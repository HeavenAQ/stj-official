package api

import (
	"database/sql"
	"net/http"
	"stj-ecommerce/utils"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	AccessToken string           `json:"access_token"`
	User        UserInfoResponse `json:"user"`
}

func (server *Server) UserLogin(ctx *gin.Context) {
	var req loginRequest
	// ensure the request is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if the user exists
	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// check if the password is correct
	err = utils.CheckPassword(user.Password, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// create access token
	accessToken, err := server.tokenMaker.CreateToken(
		user.Email,
		server.config.AccessTokenDuration,
	)

	// create response
	ctx.JSON(http.StatusOK, loginResponse{
		AccessToken: accessToken,
		User:        server.userResponse(user),
	})
}

func (server *Server) UserRegister(ctx *gin.Context) {
	// TODO: email verification
	server.CreateUser(ctx)
}
