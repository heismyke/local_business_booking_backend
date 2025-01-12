package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomReviews(t  *testing.T)Review{
	user := CreateRandomUser(t)
	business := CreateRandomBusiness(t)
	arg := CreateReviewParams{
		UserID: user.ID,
		BusinessID: business.ID,
		Rating: 5,
		Comment: "This is a very good business",
	}

	review,err := testQueries.CreateReview(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, review)

	require.Equal(t, arg.UserID, review.UserID)
	require.Equal(t, arg.BusinessID, review.BusinessID)
	require.Equal(t, arg.Rating, review.Rating)
	require.Equal(t, arg.Comment, review.Comment)
		
	require.NotZero(t, review.ID)
	require.NotZero(t, review.CreatedAt)
	return review
}

func TestCreateReviews(t *testing.T){
	createRandomReviews(t)
}

func TestGetReview(t *testing.T){
	review1 := createRandomReviews(t)
	
	review2, err := testQueries.GetReview(context.Background(), review1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, review2)

	require.Equal(t, review1.ID, review2.ID)
	require.Equal(t, review1.UserID, review2.UserID)
	require.Equal(t, review1.BusinessID, review2.BusinessID)
	require.Equal(t, review1.Rating, review2.Rating)
	require.Equal(t, review1.Comment, review2.Comment)
	require.WithinDuration(t, review1.CreatedAt, review2.CreatedAt, time.Second)
	
	require.NotZero(t, review2.ID)
	require.NotZero(t, review2.CreatedAt)
}


func TestUpdateReview(t *testing.T){
	review1 := createRandomReviews(t)
	user := CreateRandomUser(t)
	business := CreateRandomBusiness(t)
	args := UpdateReviewParams{
		ID:  review1.ID,
		UserID : user.ID,
		BusinessID: business.ID,
		Comment: review1.Comment,
	}

	err := testQueries.UpdateReview(context.Background(), args)
	require.NoError(t,err)

	updatedReview, err := testQueries.GetReview(context.Background(), review1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedReview)

	require.Equal(t, args.ID, updatedReview.ID)
	require.Equal(t, args.UserID, updatedReview.UserID)
	require.Equal(t, args.BusinessID, updatedReview.BusinessID)
	require.Equal(t, args.Rating, updatedReview.Rating)
	require.Equal(t, args.Comment, updatedReview.Comment)

	require.NotZero(t, updatedReview.ID)
	require.NotZero(t, updatedReview.CreatedAt)
}

func TestDeleteReview(t *testing.T){
	review1 := createRandomReviews(t)
	_,  err := testQueries.DeleteReview(context.Background(), review1.ID)
	require.NoError(t,err)

	review2, err := testQueries.DeleteReview(context.Background(), review1.ID)
	require.Error(t, err)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, review2)
}


func TestListReviews(t *testing.T){
	for i := 0; i < 10;i++ {
		createRandomReviews(t)
	}

	args := ListReviewsParams{
		Limit: 5,
		Offset: 5,
	}
	reviews , err := testQueries.ListReviews(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, reviews, 5)

	for _, review := range reviews{
		require.NotEmpty(t, review)
	}
}

