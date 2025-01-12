package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)


func createRandomBusinessHours(t *testing.T)BusinessHour{
	business := CreateRandomBusiness(t)

	openTime, err := time.Parse("15:04:05", "09:00:00")
	require.NoError(t,err)
	closeTime, err := time.Parse("15:04:05", "09:00:00")
	require.NoError(t, err)

	arg := CreateBusinessHoursParams{
		BusinessID: business.ID,
		DayOfWeek: "Monday",
		OpenTime: openTime ,
		CloseTime: closeTime ,
	}

	business_hour, err := testQueries.CreateBusinessHours(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, business_hour)

	require.Equal(t, arg.BusinessID, business_hour.BusinessID)
	require.Equal(t, arg.DayOfWeek, business_hour.DayOfWeek)
	require.Equal(t, arg.OpenTime, business_hour.OpenTime)
	require.Equal(t, arg.CloseTime, business_hour.CloseTime)

	require.NotZero(t, business_hour.ID)
	return business_hour	
}

func TestCreateBusinessHours(t *testing.T){
	createRandomBusinessHours(t)
}

func TestGetBusinessHour(t *testing.T){
	business_hour1 := createRandomBusinessHours(t)
	business_hour2, err := testQueries.GetBusinessHour(context.Background(), business_hour1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, business_hour2)
	require.Equal(t, business_hour1.ID, business_hour2.ID)
	require.Equal(t, business_hour1.BusinessID, business_hour2.BusinessID)
	require.Equal(t, business_hour1.DayOfWeek, business_hour2.DayOfWeek)
	require.Equal(t, business_hour1.OpenTime, business_hour2.OpenTime)
	require.Equal(t, business_hour1.CloseTime, business_hour2.CloseTime)

	require.True(t, business_hour1.OpenTime.Equal(business_hour2.OpenTime))
	require.True(t, business_hour1.CloseTime.Equal(business_hour2.CloseTime))
}

func TestUpdateBusinessHour(t *testing.T){
	business := CreateRandomBusiness(t)
	business_hour1 := createRandomBusinessHours(t)
	arg := UpdateBusinessHourParams{
		BusinessID: business.ID,
		DayOfWeek: business_hour1.DayOfWeek,
		OpenTime:  business_hour1.OpenTime,
		CloseTime: business_hour1.CloseTime,
	}

	 err := testQueries.UpdateBusinessHour(context.Background(), arg)
	 require.NoError(t,err)

	 business_hour2, err := testQueries.GetBusinessHour(context.Background(), business_hour1.ID)
	 require.NoError(t, err)
	 require.NotEmpty(t, business_hour2)

	 require.Equal(t, business_hour1.ID, business_hour2.ID)
	require.Equal(t, business_hour1.BusinessID, business_hour2.BusinessID)
	require.Equal(t, business_hour1.DayOfWeek, business_hour2.DayOfWeek)
	require.Equal(t, business_hour1.OpenTime, business_hour2.OpenTime)
	require.Equal(t, business_hour1.CloseTime, business_hour2.CloseTime)
	require.True(t, business_hour1.OpenTime.Equal(business_hour2.OpenTime))
	require.True(t, business_hour1.CloseTime.Equal(business_hour2.CloseTime))
}

func TestDeleteBusinessHour(t *testing.T){
	business_hour1 := createRandomBusinessHours(t)
	
	_, err := testQueries.DeleteBusinessHour(context.Background(), business_hour1.ID)

		require.NoError(t, err)

	business_hour2, err := testQueries.DeleteBusinessHour(context.Background(), business_hour1.ID)

	require.Error(t, err)
	require.EqualError(t,err,sql.ErrNoRows.Error())
	require.Empty(t,business_hour2)
}

func TestListBusinessHours(t *testing.T){
	for i := 0; i < 10; i++{
		createRandomBusinessHours(t)
	}
	arg := ListBusinessHoursParams{
		Limit: 5,
		Offset: 5,
	}

	business_hours, err := testQueries.ListBusinessHours(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, business_hours, 5)

	for _, business_hour := range business_hours{
		require.NotEmpty(t,business_hour)
	}
}