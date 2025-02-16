package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/heismyke/local_business_booking_app/util"
	"github.com/stretchr/testify/require"
)


func createRandomBookings(t *testing.T)Booking {
	randomDays := int(util.RandomInt(1, 30))
date := time.Now().AddDate(0, 0, randomDays).Truncate(time.Hour) // Truncate to remove time portion

	user := CreateRandomUser(t)
	business := CreateRandomBusiness(t)
	arg := CreateBookingsParams{
		UserID: user.ID,
		BusinessID: business.ID,
		Service: "Haircut",
		Date: date,
	}

	booking, err := testQueries.CreateBookings(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, booking)

	require.Equal(t, arg.UserID, booking.UserID)
	require.Equal(t, arg.BusinessID, booking.BusinessID)
	require.Equal(t, arg.Service, booking.Service)
	require.Equal(t, arg.Status, booking.Status)

	require.NotZero(t, booking.ID)
	require.NotZero(t, booking.CreatedAt)

	return booking
}

func TestCreateBookings(t *testing.T){
	createRandomBookings(t)
}


func TestGetBooking(t *testing.T){
	booking1 := createRandomBookings(t) 

	booking2, err := testQueries.GetBooking(context.Background(),  booking1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, booking2)

	require.Equal(t, booking1.UserID, booking2.UserID)
	require.Equal(t, booking1.BusinessID, booking2.BusinessID)
	require.Equal(t, booking1.Service, booking2.Service)
	require.Equal(t, booking1.Date, booking2.Date)
	require.Equal(t, booking1.Status, booking2.Status)

	require.NotZero(t, booking2.ID)
	require.NotZero(t, booking2.CreatedAt)

	require.WithinDuration(t, booking1.CreatedAt, booking2.CreatedAt, time.Second)
}

func TestUpdateBooking(t *testing.T){
	user := CreateRandomUser(t)
	business := CreateRandomBusiness(t)
	booking1 := createRandomBookings(t)
	arg := UpdateBookingParams{
		ID:  booking1.ID,
		UserID: user.ID,
		BusinessID:  business.ID,	
		Service:  "Massage",
		Date:  booking1.Date,
		Status: "Confirmed",
	}

	 err := testQueries.UpdateBooking(context.Background(), arg)
	require.NoError(t, err)

	updatedBooking, err := testQueries.GetBooking(context.Background(), booking1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedBooking)

	require.Equal(t, arg.UserID, updatedBooking.UserID)
	require.Equal(t, arg.BusinessID, updatedBooking.BusinessID)
	require.Equal(t, arg.Service, updatedBooking.Service)
	require.Equal(t, arg.Date, updatedBooking.Date)
	require.Equal(t, arg.Status, updatedBooking.Status)
	
	require.NotZero(t, updatedBooking.ID)
	require.NotZero(t, updatedBooking.CreatedAt)
}


func TestDeleteBooking(t *testing.T){
	booking1 := createRandomBookings(t)
	 
	_,err := testQueries.DeleteBooking(context.Background(), booking1.ID)
	require.NoError(t, err)

	booking2, err := testQueries.DeleteBooking(context.Background(), booking1.ID)
	require.Error(t, err)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, booking2)
}

func TestListBookings(t *testing.T){
	for i := 0; i < 10;i++{
		createRandomBookings(t)
	}
	arg := ListBookingsParams{
		Limit: 5,
		Offset: 5,
	}

	bookings, err := testQueries.ListBookings(context.Background(), arg)
	require.NoError(t,err)
	require.Len(t, bookings, 5)

	for _, booking := range bookings{
		require.NotEmpty(t, booking)
	}
}
