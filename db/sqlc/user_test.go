package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/heismyke/local_business_booking_app/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User{
		arg := CreateUserParams{
		Name:  util.RandomUser(),
		Email: util.RandomEmail(),
		Phone: util.RandomPhone(),
		Role: "admin",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Role, user.Role)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T){
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.Role, user2.Role)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	args := UpdateUserParams{
		ID:    user1.ID,
		Name:  util.RandomUser(),
		Email: util.RandomEmail(),
		Phone: util.RandomPhone(),
		Role:  user1.Role,
	}

	err := testQueries.UpdateUser(context.Background(), args)
	require.NoError(t, err)

}

func TestDeleteUser(t *testing.T){
	user1 := CreateRandomUser(t)
	_, err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t,err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t,err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListAccount(t *testing.T){
	for i := 0;  i < 10; i++{
		CreateRandomUser(t)
	}
	arg := ListUsersParams{
		Limit: 5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users{
		require.NotEmpty(t, user)
	}
}