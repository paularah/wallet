package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/paularah/wallet/pkg/db/mock"
	db "github.com/paularah/wallet/pkg/db/sqlc"
	"github.com/paularah/wallet/pkg/util"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	name          string
	walletID      int64
	buildStubs    func(store *mockdb.MockStore)
	checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func TestGetWallet(t *testing.T) {
	wallet := createTestWallet()

	testCases := []TestCase{
		{
			name:     "OK",
			walletID: wallet.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(wallet, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchWallet(t, recorder.Body, wallet)
			},
		},
		// {
		// 	name:     "NOTFOUND",
		// 	walletID: wallet.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
		// 			Times(1).
		// 			Return(db.Wallet{}, sql.ErrNoRows)

		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusNotFound, recorder.Code)
		// 	},
		// },
		// {
		// 	name:     "InvalidID",
		// 	walletID: wallet.ID,
		// 	buildStubs: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			GetWallet(gomock.Any(), gomock.Any()).
		// 			Times(0)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
		{
			name:     "InternalServerError",
			walletID: wallet.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(db.Wallet{}, sql.ErrConnDone)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/wallets/%d", testCase.walletID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(t, recorder)

		})

	}

}

func createTestWallet() db.Wallet {
	return db.Wallet{
		ID:       util.RandomID(),
		Owner:    util.RandomID(),
		Balance:  util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchWallet(t *testing.T, body *bytes.Buffer, wallet db.Wallet) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var responseWallet db.Wallet

	err = json.Unmarshal(data, &responseWallet)
	require.NoError(t, err)
	require.Equal(t, wallet, responseWallet)

}
