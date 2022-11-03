package api

import (
	"testing"

	db "github.com/paularah/wallet/pkg/db/sqlc"
	"github.com/paularah/wallet/pkg/util"
	"github.com/stretchr/testify/require"
)

func createTestUser(t *testing.T) db.User {
	password := util.RandomUsername() + util.RandomName()
	passwordHash, err := util.HashPassword(password)
	require.NoError(t, err)

	return db.User{
		ID:        util.RandomID(),
		Username:  util.RandomUsername(),
		Firstname: util.RandomName(),
		Lastname:  util.RandomName(),
		Password:  passwordHash,
	}
}

func TestCreateUser(t *testing.T) {

}

func TestLoginUserWithEmail(t *testing.T) {

}
