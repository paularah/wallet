package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/paularah/wallet/pkg/util"

	"github.com/stretchr/testify/require"
)

func createRandomWallet(t *testing.T) Wallet {
	arg := CreateWalletParams{
		Owner:    util.RandomID(),
		Balance:  util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}

	wallet, err := testQueries.CreateWallet(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, wallet)

	require.Equal(t, arg.Owner, wallet.Owner)

	require.NotZero(t, wallet.ID)
	return wallet
}
func TestCreateWallet(t *testing.T) {
	createRandomWallet(t)
}

func TestGetWallet(t *testing.T) {
	newWallet := createRandomWallet(t)
	wallet, err := testQueries.GetWallet(context.Background(), newWallet.ID)

	require.NoError(t, err)
	require.NotEmpty(t, wallet)
	require.Equal(t, newWallet.ID, wallet.ID)
	require.Equal(t, newWallet.Balance, wallet.Balance)
	require.Equal(t, newWallet.Owner, wallet.Owner)
	require.Equal(t, newWallet.Currency, wallet.Currency)
	require.WithinDuration(t, newWallet.CreatedAt, wallet.CreatedAt, time.Second)

}

func TestUpdateWallet(t *testing.T) {
	newWallet := createRandomWallet(t)

	arg := UpdateWalletParams{
		ID:      newWallet.ID,
		Balance: util.RandomAmount(),
	}

	wallet, err := testQueries.UpdateWallet(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, wallet)
	require.Equal(t, newWallet.ID, wallet.ID)
	require.Equal(t, arg.Balance, wallet.Balance)
	require.Equal(t, newWallet.Owner, wallet.Owner)
	require.Equal(t, newWallet.Currency, wallet.Currency)

}

func TestDeleteWallet(t *testing.T) {
	newWallet := createRandomWallet(t)

	err := testQueries.DeleteWallet(context.Background(), newWallet.ID)
	require.NoError(t, err)

	wallet, err := testQueries.GetWallet(context.Background(), newWallet.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, wallet)
}
