package db

import (
	"context"
	"testing"
	"tsn-ecommerce/utils"

	"github.com/stretchr/testify/require"
)

func createRandomProductTranslation(t *testing.T) ProductTranslation {
	product := createRandomProduct(t)
	args := CreateProductTranslationParams{
		ProductPk:   product.Pk,
		Language:    LanguageCode(utils.RandomLanguage()),
		Name:        utils.RandomAlphabetString(10),
		Description: utils.RandomAlphabetString(10),
		Category:    utils.RandomAlphabetString(10),
	}
	productTranslation, err := testQueries.CreateProductTranslation(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, productTranslation)
	require.NotZero(t, productTranslation.Pk)
	require.Equal(t, args.ProductPk, productTranslation.ProductPk)
	require.Equal(t, args.Language, productTranslation.Language)
	require.Equal(t, args.Name, productTranslation.Name)
	require.Equal(t, args.Description, productTranslation.Description)
	require.Equal(t, args.Category, productTranslation.Category)
	return productTranslation
}

func TestQueries_CreateProductTranslation(t *testing.T) {
	productTranslation := createRandomProductTranslation(t)
	// clean up
	testQueries.DeleteProductTranslation(context.Background(), DeleteProductTranslationParams{
		ProductPk: productTranslation.ProductPk,
		Language:  productTranslation.Language,
	})
}

func TestQueries_GetProductTranslation(t *testing.T) {
	// create a product translation and check for errors
	productTranslation1 := createRandomProductTranslation(t)
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
	require.Equal(t, productTranslation1.Description, productTranslation2.Description)
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
	productTranslation := createRandomProductTranslation(t)
	args := UpdateProductTranslationParams{
		Name:        utils.RandomAlphabetString(10),
		Description: utils.RandomAlphabetString(10),
		Category:    utils.RandomAlphabetString(10),
		Language:    productTranslation.Language,
		ProductPk:   productTranslation.ProductPk,
	}
	productTranslation, err := testQueries.UpdateProductTranslation(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, productTranslation)

	// ensure everything is updated
	require.Equal(t, args.Name, productTranslation.Name)
	require.Equal(t, args.Category, productTranslation.Category)
	require.Equal(t, args.Language, productTranslation.Language)
	require.Equal(t, args.ProductPk, productTranslation.ProductPk)
	require.Equal(t, args.Description, productTranslation.Description)
	require.WithinDuration(t, productTranslation.CreatedAt.Time, productTranslation.CreatedAt.Time, 0)

	// clean up
	testQueries.DeleteProductTranslation(context.Background(), DeleteProductTranslationParams{
		ProductPk: productTranslation.ProductPk,
		Language:  productTranslation.Language,
	})
}

func TestQueries_DeleteProductTranslation(t *testing.T) {
	// get the product translation
	productTranslation := createRandomProductTranslation(t)
	args := DeleteProductTranslationParams{
		ProductPk: productTranslation.ProductPk,
		Language:  productTranslation.Language,
	}
	testQueries.DeleteProductTranslation(context.Background(), args)

	// check if the product translation is deleted
	_, err := testQueries.GetProductTranslation(context.Background(), GetProductTranslationParams{
		ProductPk: args.ProductPk,
		Language:  args.Language,
	})

	// check if there is an error
	require.Error(t, err)
}
