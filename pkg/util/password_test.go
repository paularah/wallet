package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := "hefeifhofjwofjwfjosefjwefoejfwfoslz"

	passwordHash, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, passwordHash)

	fmt.Println(passwordHash, password)
	err = ComparePassword(password, passwordHash)

	require.NoError(t, err)

}
