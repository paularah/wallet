package db

import (
	"context"
	"testing"

	"github.com/paularah/wallet/pkg/util"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createTestUser(t)

}

func TestGetUser(t *testing.T) {
	userA := createTestUser(t)
	userB, err := testQueries.GetUser(context.Background(), userA.ID)

	require.NoError(t, err)
	require.NotEmpty(t, userB)
	require.Equal(t, userA.ID, userB.ID)
	require.Equal(t, userA.Firstname, userB.Firstname)
	require.Equal(t, userA.Email, userB.Email)

}

func createTestUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:  util.RandomUsername(),
		Firstname: util.RandomName(),
		Lastname:  util.RandomName(),
		Email:     util.RandomEmail(),
		Password:  "xxxxxxx",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Firstname, user.Firstname)
	require.Equal(t, arg.Lastname, user.Lastname)
	require.Equal(t, arg.Email, user.Email)
	return user
}
