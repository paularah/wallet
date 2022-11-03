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
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/paularah/wallet/pkg/db/mock"
	db "github.com/paularah/wallet/pkg/db/sqlc"
	"github.com/paularah/wallet/pkg/jwt"
	"github.com/paularah/wallet/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestGetWallet(t *testing.T) {
	type TestCase struct {
		name          string
		walletID      int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		setupAuth     func(t *testing.T, request *http.Request, tokener jwt.Tokener)
	}
	user := createTestUser(t)
	wallet := createTestWallet(user.ID)

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
			setupAuth: func(t *testing.T, request *http.Request, tokener jwt.Tokener) {
				jwtToken, _, err := tokener.CreateJWTToken(user.ID, time.Minute*2)
				require.NoError(t, err)
				authHeader := fmt.Sprintf("bearer %s", jwtToken)
				request.Header.Set(authHeaderKey, authHeader)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchWallet(t, recorder.Body, wallet)
			},
		},
		{
			name:     "NOTFOUND",
			walletID: wallet.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(db.Wallet{}, sql.ErrNoRows)

			},
			setupAuth: func(t *testing.T, request *http.Request, tokener jwt.Tokener) {
				jwtToken, _, err := tokener.CreateJWTToken(user.ID, time.Minute*2)
				require.NoError(t, err)
				authHeader := fmt.Sprintf("bearer %s", jwtToken)
				request.Header.Set(authHeaderKey, authHeader)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:     "InvalidID",
			walletID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokener jwt.Tokener) {
				jwtToken, _, err := tokener.CreateJWTToken(user.ID, time.Minute*2)
				require.NoError(t, err)
				authHeader := fmt.Sprintf("bearer %s", jwtToken)
				request.Header.Set(authHeaderKey, authHeader)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:     "InternalServerError",
			walletID: wallet.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetWallet(gomock.Any(), gomock.Eq(wallet.ID)).
					Times(1).
					Return(db.Wallet{}, sql.ErrConnDone)

			},
			setupAuth: func(t *testing.T, request *http.Request, tokener jwt.Tokener) {
				jwtToken, _, err := tokener.CreateJWTToken(user.ID, time.Minute*2)
				require.NoError(t, err)
				authHeader := fmt.Sprintf("bearer %s", jwtToken)
				request.Header.Set(authHeaderKey, authHeader)
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

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/wallets/%d", testCase.walletID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			testCase.setupAuth(t, request, server.tokener)

			server.router.ServeHTTP(recorder, request)

			testCase.checkResponse(t, recorder)

		})

	}

}

func TestCreateTransfer(t *testing.T) {

}

func createTestWallet(owner int64) db.Wallet {
	return db.Wallet{
		ID:       util.RandomID(),
		Owner:    owner,
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
