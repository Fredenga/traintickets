package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/fredenga/traintickets/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser (t *testing.T) User{
	password := util.RandomWords()
	hashed_password, hashErr := util.HashPassword(password)
	require.NoError(t, hashErr)
	require.NotEmpty(t, hashed_password)

	arg := CreateuserParams{
		Email: util.RandomEmail(),
		Password: hashed_password,
	}
	user, err := testQueries.Createuser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.Getuser(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
}

func TestListUser(t *testing.T){
	for i := 0; i < 5; i++ {
		createRandomUser(t)
	}
	users, err := testQueries.Listusers(context.Background(), 5)
	require.NoError(t, err)
	require.LessOrEqual(t, len(users), 5)
	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.Deleteuser(context.Background(), user1.Email)
	require.NoError(t, err)
	user2, err := testQueries.Getuser(context.Background(), user1.Email)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
