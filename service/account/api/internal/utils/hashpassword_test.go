package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswordTest(t *testing.T) {
	password := RandomString(8)

	hashPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashPassword)

	err = VerifyPassword(password, hashPassword)
	require.NoError(t, err)

	wrongPassword := RandomString(9)

	err = VerifyPassword(wrongPassword, hashPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
