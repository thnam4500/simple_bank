package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thnam4500/simple_bank/util"
)

func createRandomUser(t *testing.T) User {

	arg := CreateUserParams{
		Username:     util.RandomOwnerName(),
		HashPassword: "secret",
		FullName:     util.RandomOwnerName(),
		Email:        util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashPassword, arg.HashPassword)
	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Email, arg.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangeAt.IsZero())
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)

	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.HashPassword, user2.HashPassword)
	require.Equal(t, user1.Email, user2.Email)
}
