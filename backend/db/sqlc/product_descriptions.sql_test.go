package db

import (
	"context"
	"stj-ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomProductDescription(t *testing.T, productTranslation ProductTranslation) ProductDescription {
	// create product description
	args := CreateProductDescriptionParams{
		ProductTranslationPk: productTranslation.Pk,
		Introduction:         utils.RandomAlphabetString(10),
		Prize:                utils.RandomAlphabetString(10),
		ItemInfo:             utils.RandomAlphabetString(10),
		Recommendation:       utils.RandomAlphabetString(10),
	}
	productDescription, err := testQueries.CreateProductDescription(context.Background(), args)

	// ensure no errors
	require.NoError(t, err)
	require.NotEmpty(t, productDescription)
	require.NotZero(t, productDescription.Pk)
	require.Equal(t, args.ProductTranslationPk, productDescription.ProductTranslationPk)
	require.Equal(t, args.Introduction, productDescription.Introduction)
	require.Equal(t, args.Prize, productDescription.Prize)
	require.Equal(t, args.ItemInfo, productDescription.ItemInfo)
	require.Equal(t, args.Recommendation, productDescription.Recommendation)

	return productDescription
}

func TestQueries_GetProductDescriptionByPk(t *testing.T) {
	// add a product description
	product := createRandomProduct(t)
	productTranslation := createRandomProductTranslation(t, product)
	productDescription := createRandomProductDescription(t, productTranslation)

	// try to get it
	retrievedProductDescription, err := testQueries.GetProductDescriptionByPk(context.Background(), productDescription.Pk)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedProductDescription)
	require.Equal(t, productDescription.Pk, retrievedProductDescription.Pk)
	require.Equal(t, productDescription.ProductTranslationPk, retrievedProductDescription.ProductTranslationPk)
	require.Equal(t, productDescription.Introduction, retrievedProductDescription.Introduction)
	require.Equal(t, productDescription.Prize, retrievedProductDescription.Prize)
	require.Equal(t, productDescription.ItemInfo, retrievedProductDescription.ItemInfo)
	require.Equal(t, productDescription.Recommendation, retrievedProductDescription.Recommendation)

	// clean up
	testQueries.DeleteProductDescription(context.Background(), productDescription.Pk)
}

func TestQueries_GetProductDescriptionByTranslationPk(t *testing.T) {
	// add a product description
	product := createRandomProduct(t)
	productTranslation := createRandomProductTranslation(t, product)
	productDescription := createRandomProductDescription(t, productTranslation)

	// try to get it with productd translation pk
	retrievedProductDescription, err := testQueries.GetProductDescriptionsByTranslationPk(context.Background(), productDescription.ProductTranslationPk)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedProductDescription)
	require.Equal(t, productDescription.Pk, retrievedProductDescription.Pk)
	require.Equal(t, productDescription.ProductTranslationPk, retrievedProductDescription.ProductTranslationPk)
	require.Equal(t, productDescription.Introduction, retrievedProductDescription.Introduction)
	require.Equal(t, productDescription.Prize, retrievedProductDescription.Prize)
	require.Equal(t, productDescription.ItemInfo, retrievedProductDescription.ItemInfo)
	require.Equal(t, productDescription.Recommendation, retrievedProductDescription.Recommendation)
}

func TestQueries_CreateProductDescription(t *testing.T) {
	// create a product description
	product := createRandomProduct(t)
	productTranslation := createRandomProductTranslation(t, product)
	productDescription := createRandomProductDescription(t, productTranslation)
	testQueries.DeleteProductDescription(context.Background(), productDescription.Pk)
}

func TestQueries_DeleteProductDescription(t *testing.T) {
	// create a product description
	product := createRandomProduct(t)
	productTranslation := createRandomProductTranslation(t, product)
	productDescription := createRandomProductDescription(t, productTranslation)

	// delete the product description
	testQueries.DeleteProductDescription(context.Background(), productDescription.Pk)
}
