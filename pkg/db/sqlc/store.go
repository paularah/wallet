package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	TransferTx(ctx context.Context, arg TranferTxParams) (TransferTxResult, error)
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}

}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TranferTxParams struct {
	SenderWalletID   int64 `json:"from_wallet_id"`
	ReceiverWalletID int64 `json:"to_wallet_id"`
	Amount           int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer            Transfer `json:"transfer"`
	SenderWallet        Wallet   `json:"from_wallet"`
	ReceiverWallet      Wallet   `json:"to_wallet"`
	SenderWalletEntry   Entry    `json:"from_entry"`
	ReceiverWalletEntry Entry    `json:"to_entry"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TranferTxParams) (TransferTxResult, error) {
	var txResult TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		txResult.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			SenderWalletID:   arg.SenderWalletID,
			ReceiverWalletID: arg.ReceiverWalletID,
			Amount:           arg.Amount,
		})

		if err != nil {
			return err
		}

		txResult.SenderWalletEntry, err = q.CreateWalletEntry(ctx, CreateWalletEntryParams{
			WalletID: arg.SenderWalletID,
			Amount:   -arg.Amount,
		})

		if err != nil {
			return err
		}

		txResult.ReceiverWalletEntry, err = q.CreateWalletEntry(ctx, CreateWalletEntryParams{
			WalletID: arg.ReceiverWalletID,
			Amount:   arg.Amount,
		})

		if err != nil {
			return err
		}

		//avoids potential deadlock situation
		if arg.SenderWalletID < arg.ReceiverWalletID {
			txResult.SenderWallet, txResult.ReceiverWallet, err = addMoney(ctx, q, arg.SenderWalletID, arg.ReceiverWalletID, -arg.Amount, arg.Amount)
		} else {
			txResult.ReceiverWallet, txResult.SenderWallet, err = addMoney(ctx, q, arg.SenderWalletID, arg.ReceiverWalletID, arg.Amount, -arg.Amount)

		}

		if err != nil {
			return err
		}

		return nil
	})
	return txResult, err

}

func addMoney(ctx context.Context,
	q *Queries,
	senderWalletID int64,
	receiverWalletID int64,
	senderAmount int64,
	receiverAmount int64,
) (senderWallet Wallet, receiverWallet Wallet, err error) {

	senderWallet, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
		ID:     senderWalletID,
		Amount: senderAmount,
	})

	if err != nil {
		return
	}

	receiverWallet, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
		ID:     receiverWalletID,
		Amount: receiverAmount,
	})

	if err != nil {
		return
	}

	return

}
