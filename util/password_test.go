package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordWorksWell(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}

func TestPasswordFailsGracefully(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(RandomString(5), hashedPassword)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

func TestHashPasswordWorksWell(t *testing.T) {
	password:= RandomString(6)

	hashedPassword , err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
}

