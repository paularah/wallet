// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	AddWalletBalance(ctx context.Context, arg AddWalletBalanceParams) (Wallet, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error)
	CreateWalletEntry(ctx context.Context, arg CreateWalletEntryParams) (Entry, error)
	DeleteWallet(ctx context.Context, id int64) error
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetWallet(ctx context.Context, id int64) (Wallet, error)
	GetWalletEntry(ctx context.Context, id int64) (Entry, error)
	GetWalletForUpdate(ctx context.Context, id int64) (Wallet, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	ListWalletEntries(ctx context.Context, arg ListWalletEntriesParams) (Entry, error)
	ListWallets(ctx context.Context, arg ListWalletsParams) ([]Wallet, error)
	UpdateWallet(ctx context.Context, arg UpdateWalletParams) (Wallet, error)
}

var _ Querier = (*Queries)(nil)
