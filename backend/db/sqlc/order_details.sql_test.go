package db

import (
	"context"
	"testing"
	"tsn-ecommerce/utils"

	"github.com/stretchr/testify/require"
)

func createRandomOrderDetail(t *testing.T) OrderDetail {
	order := createRandomOrder(t)
	product := createRandomProduct(t)
	args := CreateOrderDetailParams{
		OrderPk:   order.Pk,
		ProductPk: product.Pk,
		Quantity:  1,
	}
	return createOrderDetailWithArgs(t, args)
}

func createRandomOrderDetailWithOrder(t *testing.T, order Order) OrderDetail {
	product := createRandomProduct(t)
	args := CreateOrderDetailParams{
		OrderPk:   order.Pk,
		ProductPk: product.Pk,
		Quantity:  int32(utils.RandomInt(1, 10)),
	}
	return createOrderDetailWithArgs(t, args)
}

func createOrderDetailWithArgs(t *testing.T, args CreateOrderDetailParams) OrderDetail {
	orderDetail, err := testQueries.CreateOrderDetail(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetail)
	require.NotZero(t, orderDetail.Pk)
	require.Equal(t, args.OrderPk, orderDetail.OrderPk)
	require.Equal(t, args.ProductPk, orderDetail.ProductPk)
	require.Equal(t, args.Quantity, orderDetail.Quantity)
	return orderDetail
}

// order detail creation
func TestQueries_CreateOrderDetail(t *testing.T) {
	orderDetail := createRandomOrderDetail(t)
	testQueries.DeleteOrderDetail(context.Background(), orderDetail.Pk)
}

// get order detail
func TestQueries_GetOrderDetail(t *testing.T) {
	// create an order detail and check for errors
	orderDetail1 := createRandomOrderDetail(t)
	orderDetail2, err := testQueries.GetOrderDetail(context.Background(), orderDetail1.Pk)
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
	testQueries.DeleteOrderDetail(context.Background(), orderDetail1.Pk)
}

func TestQueries_GetOrderDetailsByOrder(t *testing.T) {
	// create an order detail and check for errors
	n := 10
	order := createRandomOrder(t)
	orderDetails := make([]OrderDetail, n)
	for i := 0; i < n; i++ {
		orderDetails[i] = createRandomOrderDetailWithOrder(t, order)
	}

	// get order details by order and check for errors
	orderDetails, err := testQueries.GetOrderDetailsByOrder(context.Background(), order.Pk)
	require.NoError(t, err)
	require.Len(t, orderDetails, n)

	// check if the order details are created correctly
	for _, orderDetail := range orderDetails {
		require.NotEmpty(t, orderDetail)
		require.Equal(t, order.Pk, orderDetail.OrderPk)
		// clean up
		testQueries.DeleteOrderDetail(context.Background(), orderDetail.Pk)
		testQueries.DeleteProduct(context.Background(), orderDetail.ProductPk)
	}
}

// update order detail
func TestQueries_UpdateOrderDetail(t *testing.T) {
	// create an order detail and check for errors
	orderDetail1 := createRandomOrderDetail(t)
	product := createRandomProduct(t)
	order := createRandomOrder(t)
	args := UpdateOrderDetailParams{
		Pk:        orderDetail1.Pk,
		OrderPk:   order.Pk,
		ProductPk: product.Pk,
		Quantity:  int32(utils.RandomInt(1, 10)),
	}

	// update the order detail and check for errors
	orderDetail2, err := testQueries.UpdateOrderDetail(context.Background(), args)
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
	testQueries.DeleteOrderDetail(context.Background(), orderDetail1.Pk)
	testQueries.DeleteOrder(context.Background(), order.Pk)
	testQueries.DeleteProduct(context.Background(), product.Pk)
}

// delete order detail
func TestQueries_DeleteOrderDetail(t *testing.T) {
	// delete order detail and check for errors
	orderDetail := createRandomOrderDetail(t)
	err := testQueries.DeleteOrderDetail(context.Background(), orderDetail.Pk)
	require.NoError(t, err)

	// ensure the order detail is deleted
	_, err = testQueries.GetOrderDetail(context.Background(), orderDetail.Pk)
	require.Error(t, err)
}
