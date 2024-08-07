package db

import (
	"context"
	"stj-ecommerce/utils"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	user := createRandomUser(t)
	args := CreateOrderParams{
		UserPk:          user.Pk,
		Status:          OrderStatusPending,
		TotalPrice:      100,
		ShippingAddress: utils.RandomAlphabetString(10),
		Email:           utils.RandomAlphabetString(10),
		Phone:           utils.RandomAlphabetString(10),
		ShippingDate:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
		DeliveredDate:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	return createOrderWithArgs(t, Order{}, args)
}

func createRandomOrderWithUser(t *testing.T, user User) Order {
	args := CreateOrderParams{
		UserPk:          user.Pk,
		Status:          OrderStatusPending,
		TotalPrice:      100,
		ShippingAddress: utils.RandomAlphabetString(10),
		Email:           utils.RandomAlphabetString(10),
		Phone:           utils.RandomAlphabetString(10),
		ShippingDate:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
		DeliveredDate:   pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	return createOrderWithArgs(t, Order{}, args)
}

func createOrderWithArgs(t *testing.T, order Order, args CreateOrderParams) Order {
	// create order
	order, err := testQueries.CreateOrder(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	// check if the order is created correctly
	require.NotZero(t, order.Pk)
	require.NotZero(t, order.ID)
	require.Equal(t, args.Status, order.Status)
	require.Equal(t, args.TotalPrice, order.TotalPrice)
	require.Equal(t, args.ShippingAddress, order.ShippingAddress)
	require.Equal(t, args.Email, order.Email)
	require.Equal(t, args.Phone, order.Phone)
	require.WithinDuration(t, args.ShippingDate.Time, order.ShippingDate.Time, time.Second*3)
	require.WithinDuration(t, args.DeliveredDate.Time, order.DeliveredDate.Time, time.Second*3)
	require.NotZero(t, order.CreatedAt)
	require.NotZero(t, order.UpdatedAt)
	return order
}

func checkSameOrder(t *testing.T, order1, order2 Order) {
	require.Equal(t, order1.Pk, order2.Pk)
	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.UserPk, order2.UserPk)
	require.Equal(t, order1.Status, order2.Status)
	require.Equal(t, order1.IsPaid, order2.IsPaid)
	require.Equal(t, order1.TotalPrice, order2.TotalPrice)
	require.Equal(t, order1.ShippingAddress, order2.ShippingAddress)
	require.Equal(t, order1.Email, order2.Email)
	require.Equal(t, order1.Phone, order2.Phone)
	require.WithinDuration(t, order1.ShippingDate.Time, order2.ShippingDate.Time, 0)
	require.WithinDuration(t, order1.DeliveredDate.Time, order2.DeliveredDate.Time, 0)
	require.WithinDuration(t, order1.CreatedAt.Time, order2.CreatedAt.Time, 0)
	require.WithinDuration(t, order1.UpdatedAt.Time, order2.UpdatedAt.Time, 0)
}

// test order creation
func TestQueries_CreateOrder(t *testing.T) {
	order := createRandomOrder(t)
	testQueries.DeleteOrder(context.Background(), order.Pk)
}

func TestQueries_GetOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	// get order and check for error
	order2, err := testQueries.GetOrder(context.Background(), order1.Pk)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	// check if the orders are the same
	checkSameOrder(t, order1, order2)

	// clean up
	testQueries.DeleteOrder(context.Background(), order1.Pk)
}

func TestQueries_GetOrderByUserAndOrderID(t *testing.T) {
	// create random user for testing
	user := createRandomUser(t)
	require.NotEmpty(t, user)

	// create random order for testing
	order := createRandomOrderWithUser(t, user)
	require.NotEmpty(t, order)

	// get order by user and order id
	args := GetOrderByUserAndOrderIDParams{
		UserPk: user.Pk,
		ID:     order.ID,
	}
	order2, err := testQueries.GetOrderByUserAndOrderID(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, order2)
	require.Equal(t, order.Pk, order2.Pk)
	require.Equal(t, order.ID, order2.ID)
	require.Equal(t, order.UserPk, order2.UserPk)
	require.Equal(t, order.Status, order2.Status)
	require.Equal(t, order.IsPaid, order2.IsPaid)
	require.Equal(t, order.TotalPrice, order2.TotalPrice)
	require.Equal(t, order.ShippingAddress, order2.ShippingAddress)
	require.Equal(t, order.Email, order2.Email)
	require.Equal(t, order.Phone, order2.Phone)
	require.WithinDuration(t, order.ShippingDate.Time, order2.ShippingDate.Time, 0)
	require.WithinDuration(t, order.DeliveredDate.Time, order2.DeliveredDate.Time, 0)
	require.WithinDuration(t, order.CreatedAt.Time, order2.CreatedAt.Time, 0)
	require.WithinDuration(t, order.UpdatedAt.Time, order2.UpdatedAt.Time, 0)
}

func TestQueries_ListOrders(t *testing.T) {
	// create 10 random orders
	orders := make([]Order, 10)
	for i := 0; i < 10; i++ {
		orders[i] = createRandomOrder(t)
	}

	// list orders
	args := ListOrdersParams{
		Limit:  10,
		Offset: 0,
	}
	listedOrders, err := testQueries.ListOrders(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, listedOrders)

	// check the orders are not empty
	for _, order := range listedOrders {
		require.NotEmpty(t, order)
	}

	// clean up
	for _, order := range orders {
		testQueries.DeleteOrder(context.Background(), order.Pk)
	}
}

func TestQueries_ListOrdersByUser(t *testing.T) {
	user := createRandomUser(t)
	orders := make([]Order, 10)
	for i := 0; i < 10; i++ {
		orders[i] = createRandomOrderWithUser(t, user)
	}

	// list orders
	args := ListOrdersByUserParams{
		UserPk: user.Pk,
		Limit:  10,
		Offset: 0,
	}
	listedOrders, err := testQueries.ListOrdersByUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, listedOrders)

	// check the orders are not empty
	for _, order := range listedOrders {
		require.NotEmpty(t, order)
	}

	// check if user is the same
	for _, order := range listedOrders {
		require.Equal(t, user.Pk, order.UserPk)
	}

	// clean up
	testQueries.DeleteUser(context.Background(), user.Pk)
	for _, order := range orders {
		testQueries.DeleteOrder(context.Background(), order.Pk)
	}
}

func TestQueries_ListOrdersByStatus(t *testing.T) {
	orders := make([]Order, 10)
	for i := 0; i < 10; i++ {
		orders[i] = createRandomOrder(t)
	}

	// list orders
	args := ListOrdersByStatusParams{
		Status: OrderStatusPending,
		Limit:  10,
		Offset: 0,
	}
	listedOrders, err := testQueries.ListOrdersByStatus(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, listedOrders)

	// check the orders are not empty
	for _, order := range listedOrders {
		require.NotEmpty(t, order)
	}

	// check if status is the same
	for _, order := range listedOrders {
		require.Equal(t, OrderStatusPending, order.Status)
	}

	// clean up
	for _, order := range orders {
		testQueries.DeleteOrder(context.Background(), order.Pk)
	}
}

// Update order
func TestQueries_UpdateOrder(t *testing.T) {
	user := createRandomUser(t)
	order1 := createRandomOrder(t)
	args := UpdateOrderParams{
		Pk:              order1.Pk,
		UserPk:          user.Pk,
		Status:          OrderStatusDelivered,
		TotalPrice:      int32(utils.RandomInt(100, 1000)),
		ShippingAddress: utils.RandomAlphabetString(10),
		Email:           utils.RandomAlphabetString(10),
		Phone:           utils.RandomAlphabetString(10),
		ShippingDate:    order1.ShippingDate,
		DeliveredDate:   order1.DeliveredDate,
		IsPaid:          true,
	}
	order2, err := testQueries.UpdateOrder(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	// check if the orders are the same
	require.Equal(t, order1.Pk, order2.Pk)
	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order2.UserPk, args.UserPk)
	require.Equal(t, args.Status, order2.Status)
	require.Equal(t, order2.IsPaid, true)
	require.Equal(t, args.TotalPrice, order2.TotalPrice)
	require.Equal(t, args.ShippingAddress, order2.ShippingAddress)
	require.Equal(t, args.Email, order2.Email)
	require.Equal(t, args.Phone, order2.Phone)
	require.WithinDuration(t, order1.ShippingDate.Time, order2.ShippingDate.Time, 0)
	require.WithinDuration(t, order1.DeliveredDate.Time, order2.DeliveredDate.Time, 0)
	require.WithinDuration(t, order1.CreatedAt.Time, order2.CreatedAt.Time, 0)
	require.NotEqual(t, order1.UpdatedAt.Time, order2.UpdatedAt.Time)

	// clean up
	testQueries.DeleteUser(context.Background(), order1.UserPk)
	testQueries.DeleteOrder(context.Background(), order1.Pk)
	testQueries.DeleteUser(context.Background(), user.Pk)
}

// test order deletion
func TestQueries_DeleteOrder(t *testing.T) {
	order := createRandomOrder(t)
	err := testQueries.DeleteOrder(context.Background(), order.Pk)
	require.NoError(t, err)

	// check if the order is deleted correctly
	order, err = testQueries.GetOrder(context.Background(), order.Pk)
	require.Error(t, err)
}
