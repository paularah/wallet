package db

import (
	"context"
	"testing"

	"github.com/paularah/wallet/pkg/util"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	CreateTestUser(t)

}

func TestGetUser(t *testing.T) {
	userA := CreateTestUser(t)
	userB, err := testQueries.GetUser(context.Background(), userA.ID)

	require.NoError(t, err)
	require.NotEmpty(t, userB)
	require.Equal(t, userA.ID, userB.ID)

}

func CreateTestUser(t *testing.T) User {
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
