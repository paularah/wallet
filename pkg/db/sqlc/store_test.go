package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	walletA := createRandomWallet(t)
	walletB := createRandomWallet(t)

	n := 6
	amount := int64(30)

	txErrors := make(chan error)
	txResults := make(chan TransferTxResult)

	for i := 0; i <= n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TranferTxParams{
				SenderWalletID:   walletA.ID,
				ReceiverWalletID: walletB.ID,
				Amount:           amount,
			})

			txErrors <- err
			txResults <- result

		}()
	}

	for i := 0; i < n; i++ {

		err := <-txErrors
		require.NoError(t, err)

		txResult := <-txResults
		require.NotEmpty(t, txResult)

		transfer := txResult.Transfer

		//check that the transfer for tx exists
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// validate the tranfer fields
		require.NotEmpty(t, transfer)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		require.Equal(t, transfer.Amount, amount)
		require.Equal(t, walletA.ID, transfer.SenderWalletID)
		require.Equal(t, walletB.ID, transfer.ReceiverWalletID)

		//check the from entry \for the tx exixst
		_, err = store.GetWalletEntry(context.Background(), txResult.SenderWalletEntry.ID)
		require.NoError(t, err)

		senderWalletEntry := txResult.SenderWalletEntry

		require.NotEmpty(t, senderWalletEntry)
		require.NotZero(t, senderWalletEntry.ID)
		require.NotZero(t, senderWalletEntry.CreatedAt)
		require.Equal(t, senderWalletEntry.WalletID, walletA.ID)
		require.Equal(t, senderWalletEntry.Amount, -amount)

		//check the from entry \for the tx exists
		_, err = store.GetWalletEntry(context.Background(), txResult.ReceiverWalletEntry.ID)
		require.NoError(t, err)

		receiverWalletEntry := txResult.ReceiverWalletEntry

		require.NotEmpty(t, receiverWalletEntry)
		require.NotZero(t, receiverWalletEntry.ID)
		require.NotZero(t, receiverWalletEntry.CreatedAt)
		require.Equal(t, receiverWalletEntry.WalletID, walletB.ID)
		require.Equal(t, receiverWalletEntry.Amount, amount)

		//check wallets
		senderWallet := txResult.SenderWallet
		require.NotEmpty(t, senderWallet)
		require.Equal(t, senderWallet.ID, walletA.ID)

		receiverWallet := txResult.ReceiverWallet
		require.NotEmpty(t, receiverWallet)
		require.Equal(t, receiverWallet.ID, walletB.ID)

		//diifs between wallet balances after tx
		diffA := walletA.Balance - senderWallet.Balance
		diffB := receiverWallet.Balance - walletB.Balance

		fmt.Println(walletB, receiverWallet)

		require.True(t, diffA > 0 && diffB > 0)
		require.Equal(t, diffA, diffB)
		require.True(t, diffA%amount == 0)

		x := int(diffA / amount)
		require.True(t, x >= 1 && x <= n)
	}

	//check that balances n the wallet have been updated

	updatedWalletA, err := testQueries.GetWallet(context.Background(), walletA.ID)
	require.NoError(t, err)

	updatedWalletB, err := testQueries.GetWallet(context.Background(), walletB.ID)
	require.NoError(t, err)

	require.Equal(t, updatedWalletA.Balance, walletA.Balance-amount*int64(n))
	require.Equal(t, updatedWalletB.Balance, walletB.Balance+amount*int64(n))

}
