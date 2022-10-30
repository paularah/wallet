package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/paularah/wallet/pkg/db/sqlc"
)

type createTransferRequest struct {
	SenderWalletID   int64  `json:"sender_wallet_id" binding:"required,min=1"`
	ReceiverWalletID int64  `json:"receiver_wallet_id" binding:"required,min=1"`
	Amount           int64  `json:"amount" binding:"required,gt=0"`
	Currency         string `json:"currency" binding:"required,oneof=RWF"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TranferTxParams{
		SenderWalletID:   req.SenderWalletID,
		ReceiverWalletID: req.ReceiverWalletID,
		Amount:           req.Amount,
	}

	isValid := server.validateTransfer(ctx, req)

	if !isValid {
		return
	}

	transfer, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}

func (server *Server) validateTransfer(ctx *gin.Context, arg createTransferRequest) bool {
	senderWallet, err := server.store.GetWallet(ctx, arg.SenderWalletID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	receiverWallet, err := server.store.GetWallet(ctx, arg.ReceiverWalletID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if senderWallet.Balance < arg.Amount {
		err := fmt.Errorf("insufficient balance")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	if senderWallet.Currency != receiverWallet.Currency ||
		senderWallet.Currency != arg.Currency {
		err := fmt.Errorf("currency mismatch")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
