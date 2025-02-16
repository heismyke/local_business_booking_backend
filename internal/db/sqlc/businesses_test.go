package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/heismyke/local_business_booking_app/util"
	"github.com/stretchr/testify/require"
)
func CreateRandomBusiness(t *testing.T) Business{
	user := CreateRandomUser(t) 
	servicesJSON, err := json.Marshal(map[string]string{
		"service1": "Software Development",
		"service2": "Cloud Hosting",
		"service3": "IT Consulting",
	})
	require.NoError(t, err)
	arg := CreateBusinessesParams{
		Owner: user.ID,
		Name: "Elite Tech Solutions",
		Address: "12 Silicon Valley, San Francisco, CA",
		Lattitude:  37.7749,
		Longitude: -122.4194,
		Phone: util.RandomPhone(),
		Email: util.RandomEmail(),
		Category: "Tech",
		Services: servicesJSON,
	}

	business, err := testQueries.CreateBusinesses(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, business)

	require.Equal(t, arg.Owner, business.Owner)
	require.Equal(t, arg.Name, business.Name)
	require.Equal(t, arg.Address, business.Address)
	require.Equal(t, arg.Lattitude, business.Lattitude)
	require.Equal(t, arg.Longitude, business.Longitude)
	require.Equal(t, arg.Phone, business.Phone)
	require.Equal(t, arg.Email, business.Email)
	require.Equal(t, arg.Category, business.Category)
	var expectedServices, actualServices map[string]string
	require.NoError(t, json.Unmarshal(arg.Services, &expectedServices))
	require.NoError(t, json.Unmarshal(business.Services, &actualServices))
	require.Equal(t, expectedServices, actualServices)

	require.NotZero(t, business.ID)
	require.NotZero(t, business.CreatedAt)

	return business
}

func TestCreateBusinesses(t *testing.T){
	CreateRandomBusiness(t)
}

func TestGetBusiness(t *testing.T){
	business1 := CreateRandomBusiness(t)
	business2, err := testQueries.GetBusinesses(context.Background(), business1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, business2)
	require.Equal(t, business1.ID, business2.ID)
	require.Equal(t, business1.Owner, business2.Owner)
	require.Equal(t, business1.Name, business2.Name)
	require.Equal(t, business1.Address, business2.Address)
	require.Equal(t, business1.Lattitude, business2.Lattitude)
	require.Equal(t, business1.Phone, business2.Phone)
	require.Equal(t, business1.Email, business2.Email)
	require.Equal(t, business1.Category, business2.Category)
	require.Equal(t, business1.Services, business2.Services)
	require.WithinDuration(t, business1.CreatedAt, business2.CreatedAt, time.Second)
	
}


func TestUpdateBusiness(t *testing.T){
	user := CreateRandomUser(t)
	business1 := CreateRandomBusiness(t)
	args := UpdateBusinessParams{
		ID:  business1.ID,
		Owner: user.ID,
		Name: business1.Name,
		Address: business1.Address,
		Lattitude:  business1.Lattitude,
		Longitude: business1.Longitude,
		Phone: util.RandomPhone(),
		Email: util.RandomEmail(),
		Category: business1.Category,
		Services: business1.Services,
	}

	err := testQueries.UpdateBusiness(context.Background(), args)
	require.NoError(t, err)

	updatedBusiness, err := testQueries.GetBusinesses(context.Background(), business1.ID)
	require.NoError(t,err)
	require.NotEmpty(t, updatedBusiness)

	require.Equal(t, args.ID, updatedBusiness.ID)
	require.Equal(t, args.Owner, updatedBusiness.Owner)
	require.Equal(t, args.Name, updatedBusiness.Name)
	require.Equal(t, args.Address, updatedBusiness.Address)
	require.Equal(t, args.Lattitude, updatedBusiness.Lattitude)
	require.Equal(t, args.Longitude, updatedBusiness.Longitude)
	require.Equal(t, args.Phone, updatedBusiness.Phone)
	require.Equal(t, args.Email, updatedBusiness.Email)
	require.Equal(t, args.Category, updatedBusiness.Category)
	require.Equal(t, args.Services, updatedBusiness.Services)
}


func TestDeleteBusiness(t *testing.T){
	business := CreateRandomBusiness(t)
	_, err := testQueries.DeleteBusiness(context.Background(), business.ID)
	require.NoError(t, err)

	business2, err := testQueries.DeleteBusiness(context.Background(), business.ID)
	require.Error(t, err)	
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, business2)
}

func TestListBusinesses(t *testing.T){
	for i := 0; i < 10; i++{
		CreateRandomBusiness(t)
	}

	arg := ListBusinessesParams{
		Limit: 5,
		Offset: 5,
	}

	businesses, err := testQueries.ListBusinesses(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, businesses, 5)


	for _, business := range businesses{
		require.NotEmpty(t, business)
	}
}
