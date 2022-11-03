package jwt

import (
	"testing"
	"time"

	"github.com/paularah/wallet/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestCreateJWT(t *testing.T) {
	tokener := &Tokener{secretKey: "cmrcfejfdui23ur923oncn3nf"}
	userID := util.RandomID()
	duration := time.Minute
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	token, claim, err := tokener.CreateJWTToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, claim)

	claim, err = tokener.VerifyJWT(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, claim.ID)
	require.Equal(t, userID, claim.UserID)
	require.WithinDuration(t, issuedAt, claim.IssuedAt, time.Second)
	require.WithinDuration(t, expiresAt, claim.ExpiresAt, time.Second)
}

func TestValidClaim(t *testing.T) {

}

func TestInvalidClaim(t *testing.T) {

}

func TestInvalidJWT(T *testing.T) {

}

func TestExpiredJWT(t *testing.T) {

}
