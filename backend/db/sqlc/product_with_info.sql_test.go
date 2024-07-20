package db

import (
	"context"
	"stj-ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueries_GetProductWithInfo(t *testing.T) {
	// add a product
	productWithInfo := addRandomProductTx(t)
	args := GetProductWithInfoParams{
		ID:       productWithInfo.Product.ID,
		Language: productWithInfo.ProductTrans.Language,
	}

	// get product with info and check for error
	product2, err := testQueries.GetProductWithInfo(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	// check if the product is the same
	require.Equal(t, productWithInfo.Product.ID, product2.ID)
	require.Equal(t, productWithInfo.ProductTrans.Name, product2.Name)
	require.Equal(t, productWithInfo.ProductTrans.CreatedAt, product2.CreatedAt)
	require.Equal(t, productWithInfo.ProductTrans.UpdatedAt, product2.UpdatedAt)
	require.Equal(t, productWithInfo.ProductTrans.Category, product2.Category)
}

func TestQueries_ListProductWithInfo(t *testing.T) {
	num := 10
	langCode := LanguageCode(utils.RandomLanguage())
	// create dummy products
	for i := 0; i < num; i++ {
		addProductTxWithLangCode(t, langCode)
	}

	// list products with info and check for errors
	args := ListProductWithInfoParams{
		Language: langCode,
		Limit:    int32(num),
		Offset:   0,
	}
	products, err := testQueries.ListProductWithInfo(context.Background(), args)

	// check for errors or empty products
	require.NoError(t, err)
	for _, product := range products {
		require.NotEmpty(t, product)
	}
}
