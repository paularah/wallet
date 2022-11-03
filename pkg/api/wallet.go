package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/paularah/wallet/pkg/db/sqlc"
	"github.com/paularah/wallet/pkg/jwt"
)

type createWalletRequest struct {
	Currency string `json:"currency" binding:"required,oneof=RWF"`
}

func (server *Server) createWallet(ctx *gin.Context) {
	var req createWalletRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authClaim := ctx.MustGet(claimKey).(*jwt.Claim)

	arg := db.CreateWalletParams{
		Owner:    authClaim.UserID,
		Currency: req.Currency,
		Balance:  0,
	}

	wallet, err := server.store.CreateWallet(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if errName := pqErr.Code.Name(); errName == "foreign_key_violation" ||
				errName == "unique_violation" {
				ctx.JSON(http.StatusUnprocessableEntity, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, wallet)
}

type getWalletRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getWallet(ctx *gin.Context) {

	var req getWalletRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	wallet, err := server.store.GetWallet(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authClaim := ctx.MustGet(claimKey).(*jwt.Claim)
	if authClaim.UserID != wallet.Owner {
		err := errors.New("wallet doensn't belong to user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, wallet)

}

type listWalletRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *Server) listWallet(ctx *gin.Context) {

	var req listWalletRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authClaim := ctx.MustGet(claimKey).(*jwt.Claim)

	arg := db.ListWalletsParams{
		Owner:  authClaim.UserID,
		Limit:  int32(req.PageID),
		Offset: (int32(req.PageID) - 1) * int32(req.PageSize),
	}

	wallets, err := server.store.ListWallets(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, wallets)

}
