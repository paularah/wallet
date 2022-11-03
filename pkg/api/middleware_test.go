package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paularah/wallet/pkg/jwt"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {
	type TestCase struct {
		name         string
		setupAuth    func(t *testing.T, request *http.Request, tokener jwt.Tokener)
		checkReponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}

	testcases := []TestCase{
		{
			name: "Valid Auth Token",
			setupAuth: func(t *testing.T, request *http.Request, tokener jwt.Tokener) {
				jwtToken, _, err := tokener.CreateJWTToken(int64(4), time.Minute*2)
				require.NoError(t, err)
				authHeader := fmt.Sprintf("bearer %s", jwtToken)
				request.Header.Set(authHeaderKey, authHeader)
			},
			checkReponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusOK)
			},
		},
		{
			name: "Missing Auth",
			setupAuth: func(t *testing.T, request *http.Request, tokener jwt.Tokener) {
			},
			checkReponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, recorder.Code, http.StatusUnauthorized)
			},
		},
	}

	for i := range testcases {
		testcase := testcases[i]
		t.Run(testcase.name, func(t *testing.T) {
			server := NewTestServer(t, nil)
			path := "/auth"
			tokener := server.tokener
			server.router.GET(
				path,
				authMiddleware(tokener),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, path, nil)
			require.NoError(t, err)

			testcase.setupAuth(t, request, tokener)
			server.router.ServeHTTP(recorder, request)
			testcase.checkReponse(t, recorder)
		})
	}
}
