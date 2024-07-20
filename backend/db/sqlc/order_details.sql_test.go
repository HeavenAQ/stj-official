package db

import (
	"context"
	"stj-ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOrderDetails(t *testing.T) OrderDetail {
	order := createRandomOrder(t)
	product := createRandomProduct(t)
	args := CreateOrderDetailsParams{
		OrderPk:   order.Pk,
		ProductPk: product.Pk,
		Quantity:  1,
	}
	return createOrderDetailsWithArgs(t, args)
}

func createRandomOrderDetailsWithOrder(t *testing.T, order Order) OrderDetail {
	product := createRandomProduct(t)
	args := CreateOrderDetailsParams{
		OrderPk:   order.Pk,
		ProductPk: product.Pk,
		Quantity:  int32(utils.RandomInt(1, 10)),
	}
	return createOrderDetailsWithArgs(t, args)
}

func createOrderDetailsWithArgs(t *testing.T, args CreateOrderDetailsParams) OrderDetail {
	orderDetail, err := testQueries.CreateOrderDetails(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetail)
	require.NotZero(t, orderDetail.Pk)
	require.Equal(t, args.OrderPk, orderDetail.OrderPk)
	require.Equal(t, args.ProductPk, orderDetail.ProductPk)
	require.Equal(t, args.Quantity, orderDetail.Quantity)
	return orderDetail
}

// order detail creation
func TestQueries_CreateOrderDetails(t *testing.T) {
	orderDetail := createRandomOrderDetails(t)
	testQueries.DeleteOrderDetails(context.Background(), orderDetail.Pk)
}

// get order detail
func TestQueries_GetOrderDetails(t *testing.T) {
	// create an order detail and check for errors
	orderDetail1 := createRandomOrderDetails(t)
	orderDetail2, err := testQueries.GetOrderDetails(context.Background(), orderDetail1.Pk)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetail2)

	// check if the order detail is created correctly
	require.Equal(t, orderDetail1.Pk, orderDetail2.Pk)
	require.Equal(t, orderDetail1.OrderPk, orderDetail2.OrderPk)
	require.Equal(t, orderDetail1.ProductPk, orderDetail2.ProductPk)
	require.Equal(t, orderDetail1.Quantity, orderDetail2.Quantity)
	require.NotZero(t, orderDetail2.CreatedAt)
	require.NotZero(t, orderDetail2.UpdatedAt)

	// clean up
	testQueries.DeleteOrderDetails(context.Background(), orderDetail1.Pk)
}

func TestQueries_GetOrderDetailssByOrder(t *testing.T) {
	// create an order detail and check for errors
	n := 10
	order := createRandomOrder(t)
	orderDetails := make([]OrderDetail, n)
	for i := 0; i < n; i++ {
		orderDetails[i] = createRandomOrderDetailsWithOrder(t, order)
	}

	// get order details by order and check for errors
	orderDetails, err := testQueries.GetOrderDetailsByOrderPk(context.Background(), order.Pk)
	require.NoError(t, err)
	require.Len(t, orderDetails, n)

	// check if the order details are created correctly
	for _, orderDetail := range orderDetails {
		require.NotEmpty(t, orderDetail)
		require.Equal(t, order.Pk, orderDetail.OrderPk)
		// clean up
		testQueries.DeleteOrderDetails(context.Background(), orderDetail.Pk)
		testQueries.DeleteProduct(context.Background(), orderDetail.ProductPk)
	}
}

// update order detail
func TestQueries_UpdateOrderDetails(t *testing.T) {
	// create an order detail and check for errors
	orderDetail1 := createRandomOrderDetails(t)
	product := createRandomProduct(t)
	order := createRandomOrder(t)
	args := UpdateOrderDetailsParams{
		Pk:        orderDetail1.Pk,
		OrderPk:   order.Pk,
		ProductPk: product.Pk,
		Quantity:  int32(utils.RandomInt(1, 10)),
	}

	// update the order detail and check for errors
	orderDetail2, err := testQueries.UpdateOrderDetails(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetail2)

	// check if the order detail is updated correctly
	require.Equal(t, orderDetail1.Pk, orderDetail2.Pk)
	require.Equal(t, args.OrderPk, orderDetail2.OrderPk)
	require.Equal(t, args.ProductPk, orderDetail2.ProductPk)
	require.Equal(t, args.Quantity, orderDetail2.Quantity)
	require.WithinDuration(t, orderDetail1.CreatedAt.Time, orderDetail2.CreatedAt.Time, 0)
	require.NotEqual(t, orderDetail1.UpdatedAt.Time, orderDetail2.UpdatedAt.Time)

	// clean up
	testQueries.DeleteOrderDetails(context.Background(), orderDetail1.Pk)
	testQueries.DeleteOrder(context.Background(), order.Pk)
	testQueries.DeleteProduct(context.Background(), product.Pk)
}

// delete order detail
func TestQueries_DeleteOrderDetails(t *testing.T) {
	// delete order detail and check for errors
	orderDetail := createRandomOrderDetails(t)
	err := testQueries.DeleteOrderDetails(context.Background(), orderDetail.Pk)
	require.NoError(t, err)

	// ensure the order detail is deleted
	_, err = testQueries.GetOrderDetails(context.Background(), orderDetail.Pk)
	require.Error(t, err)
}
