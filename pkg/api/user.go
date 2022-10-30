package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/paularah/wallet/pkg/db/sqlc"
	"github.com/paularah/wallet/pkg/util"
)

type createUserRequest struct {
	Username  string `json:"username" binding:"required,alphanum"`
	Firstname string `json:"firstname" binding:"required,alphanum"`
	Lastname  string `json:"lastname" binding:"required,alphanum"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type CreateUserResponse struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Password          string    `json:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	passwordHash, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:     req.Email,
		Username:  req.Username,
		Password:  passwordHash,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if errName := pqErr.Code.Name(); errName == "unique_violation" {
				ctx.JSON(http.StatusUnprocessableEntity, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
