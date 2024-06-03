package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) Product {
	args := CreateProductParams{
		Price:     100,
		ImageURLs: []string{"image"},
		Status:    ProductStatusInStock,
		Quantity:  10,
	}

	// create a product with random data and check for errors
	product, err := testQueries.CreateProduct(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	// check if the product is created correctly
	require.NotZero(t, product.Pk)
	require.NotZero(t, product.ID)
	require.Equal(t, args.Price, product.Price)
	require.Equal(t, args.ImageURLs, product.ImageURLs)
	require.Equal(t, args.Status, product.Status)
	require.Equal(t, args.Quantity, product.Quantity)
	require.NotZero(t, product.CreatedAt)
	require.NotZero(t, product.UpdatedAt)
	return product
}

func TestQueries_CreateProduct(t *testing.T) {
	product := createRandomProduct(t)
	testQueries.DeleteProduct(context.Background(), product.Pk)
}

func TestQueries_GetProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	// get the product and check for errors
	product2, err := testQueries.GetProduct(context.Background(), product1.Pk)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	// check if the product is created correctly
	require.Equal(t, product1.Pk, product2.Pk)
	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.ImageURLs, product2.ImageURLs)
	require.Equal(t, product1.Status, product2.Status)
	require.Equal(t, product1.Quantity, product2.Quantity)
	require.WithinDuration(t, product1.CreatedAt.Time, product2.CreatedAt.Time, 0)
	require.WithinDuration(t, product1.UpdatedAt.Time, product2.UpdatedAt.Time, 0)
	require.NotZero(t, product2.CreatedAt)
	require.NotZero(t, product2.UpdatedAt)

	// clean up
	testQueries.DeleteProduct(context.Background(), product1.Pk)
}

func TestQueries_ListProducts(t *testing.T) {
	// create products
	n := 10
	for i := 0; i < n; i++ {
		createRandomProduct(t)
	}

	// list products and check for errors
	products, err := testQueries.ListProducts(context.Background(), ListProductsParams{
		Limit:  int32(n),
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, products, 10)

	// ensure no empty products
	for _, product := range products {
		require.NotEmpty(t, product)
	}

	// clean up
	for _, product := range products {
		testQueries.DeleteProduct(context.Background(), product.Pk)
	}
}

func TestQueries_UpdateProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	args := UpdateProductParams{
		Pk:        product1.Pk,
		Price:     200,
		ImageURLs: []string{"image2", "image3"},
		Status:    ProductStatusOutOfStock,
		Quantity:  20,
	}

	// update the product and check for errors
	product2, err := testQueries.UpdateProduct(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	// check if the product is updated correctly
	require.Equal(t, product1.Pk, product2.Pk)
	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, args.Price, product2.Price)
	require.Equal(t, args.ImageURLs, product2.ImageURLs)
	require.Equal(t, args.Status, product2.Status)
	require.Equal(t, args.Quantity, product2.Quantity)
	require.WithinDuration(t, product1.CreatedAt.Time, product2.CreatedAt.Time, 0)
	require.NotEqual(t, product1.UpdatedAt.Time, product2.UpdatedAt.Time)

	// clean up
	testQueries.DeleteProduct(context.Background(), product1.Pk)
}

func TestQueries_DeleteProduct(t *testing.T) {
	product := createRandomProduct(t)

	// delete the product and check for errors
	err := testQueries.DeleteProduct(context.Background(), product.Pk)
	require.NoError(t, err)

	// check if the product is deleted
	_, err = testQueries.GetProduct(context.Background(), product.Pk)
	require.Error(t, err)
}

func TestQueries_DeleteProductAndInfo(t *testing.T) {
	// create a random product with info
	product := addRandomProductTx(t)
	err := testQueries.DeleteProductById(context.Background(), product.Product.ID)
	require.NoError(t, err)

	// ensure product is deleted
	product2, err := testQueries.GetProduct(context.Background(), product.Product.Pk)
	require.Error(t, err)
	require.Empty(t, product2)

	// ensure product info is deleted
	var nilProductInfo []ProductTranslation
	product2Info, err := testQueries.GetProductTranslations(context.Background(), product.Product.Pk)
	require.NoError(t, err)
	require.Equal(t, product2Info, nilProductInfo)
}

func TestQueries_GetProductByID(t *testing.T) {
	product := createRandomProduct(t)

	// get the product and check for errors
	product2, err := testQueries.GetProductByID(context.Background(), product.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	// check if the product is created correctly
	require.Equal(t, product.Pk, product2.Pk)
	require.Equal(t, product.ID, product2.ID)
	require.Equal(t, product.Price, product2.Price)
	require.Equal(t, product.ImageURLs, product2.ImageURLs)
	require.Equal(t, product.Status, product2.Status)
	require.Equal(t, product.Quantity, product2.Quantity)
	require.WithinDuration(t, product.CreatedAt.Time, product2.CreatedAt.Time, 0)
	require.WithinDuration(t, product.UpdatedAt.Time, product2.UpdatedAt.Time, 0)
	require.NotZero(t, product2.CreatedAt)
	require.NotZero(t, product2.UpdatedAt)

	// clean up
	testQueries.DeleteProduct(context.Background(), product.Pk)
}
