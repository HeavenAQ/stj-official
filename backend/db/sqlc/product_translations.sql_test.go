package db

import (
	"context"
	"stj-ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomProductTranslation(t *testing.T, product Product) ProductTranslation {
	args := CreateProductTranslationParams{
		ProductPk: product.Pk,
		Language:  LanguageCode(utils.RandomLanguage()),
		Name:      utils.RandomAlphabetString(10),
		Category:  utils.RandomAlphabetString(10),
	}
	productTranslation, err := testQueries.CreateProductTranslation(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, productTranslation)
	require.NotZero(t, productTranslation.Pk)
	require.Equal(t, args.ProductPk, productTranslation.ProductPk)
	require.Equal(t, args.Language, productTranslation.Language)
	require.Equal(t, args.Name, productTranslation.Name)
	require.Equal(t, args.Category, productTranslation.Category)
	return productTranslation
}

func TestQueries_CreateProductTranslation(t *testing.T) {
	// create a product translation
	product := createRandomProduct(t)
	productTranslation := createRandomProductTranslation(t, product)
	// clean up
	testQueries.DeleteProductTranslation(context.Background(), DeleteProductTranslationParams{
		ProductPk: productTranslation.ProductPk,
		Language:  productTranslation.Language,
	})
}

func TestQueries_GetProductTranslation(t *testing.T) {
	// create a product translation and check for errors
	product := createRandomProduct(t)
	productTranslation1 := createRandomProductTranslation(t, product)
	productTranslation2, err := testQueries.GetProductTranslation(context.Background(), GetProductTranslationParams{
		ProductPk: productTranslation1.ProductPk,
		Language:  productTranslation1.Language,
	})
	require.NoError(t, err)
	require.NotEmpty(t, productTranslation2)

	// check if the product translation is created correctly
	require.Equal(t, productTranslation1.Pk, productTranslation2.Pk)
	require.Equal(t, productTranslation1.ProductPk, productTranslation2.ProductPk)
	require.Equal(t, productTranslation1.Language, productTranslation2.Language)
	require.Equal(t, productTranslation1.Name, productTranslation2.Name)
	require.Equal(t, productTranslation1.Category, productTranslation2.Category)
	require.WithinDuration(t, productTranslation1.CreatedAt.Time, productTranslation2.CreatedAt.Time, 0)
	require.WithinDuration(t, productTranslation1.UpdatedAt.Time, productTranslation2.UpdatedAt.Time, 0)

	// clean up
	testQueries.DeleteProductTranslation(context.Background(), DeleteProductTranslationParams{
		ProductPk: productTranslation1.ProductPk,
		Language:  productTranslation1.Language,
	})
}

func TestQueries_UpdateProductTranslation(t *testing.T) {
	// update the product translation and check for errors
	product := createRandomProduct(t)
	productTranslation := createRandomProductTranslation(t, product)
	args := UpdateProductTranslationParams{
		Name:      utils.RandomAlphabetString(10),
		Category:  utils.RandomAlphabetString(10),
		Language:  productTranslation.Language,
		ProductPk: productTranslation.ProductPk,
	}
	productTranslation, err := testQueries.UpdateProductTranslation(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, productTranslation)

	// ensure everything is updated
	require.Equal(t, args.Name, productTranslation.Name)
	require.Equal(t, args.Category, productTranslation.Category)
	require.Equal(t, args.Language, productTranslation.Language)
	require.Equal(t, args.ProductPk, productTranslation.ProductPk)
	require.WithinDuration(t, productTranslation.CreatedAt.Time, productTranslation.CreatedAt.Time, 0)

	// clean up
	testQueries.DeleteProductTranslation(context.Background(), DeleteProductTranslationParams{
		ProductPk: productTranslation.ProductPk,
		Language:  productTranslation.Language,
	})
}

func TestQueries_DeleteProductTranslation(t *testing.T) {
	// get the product translation
	product := createRandomProduct(t)
	productTranslation := createRandomProductTranslation(t, product)
	args := DeleteProductTranslationParams{
		ProductPk: productTranslation.ProductPk,
		Language:  productTranslation.Language,
	}
	testQueries.DeleteProductTranslation(context.Background(), args)

	// check if the product translation is deleted
	_, err := testQueries.GetProductTranslation(context.Background(), GetProductTranslationParams(args))

	// check if there is an error
	require.Error(t, err)
}
