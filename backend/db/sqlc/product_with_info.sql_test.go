package db

import (
	"context"
	"testing"
	"tsn-ecommerce/utils"

	"github.com/stretchr/testify/require"
)

func TestQueries_GetProductWithInfo(t *testing.T) {
	// add a product
	product := addRandomProductTx(t)
	arg := GetProductWithInfoParams{
		ID:       product.Product.ID,
		Language: product.ProductTrans.Language,
	}

	// get product with info and check for error
	product2, err := testQueries.GetProductWithInfo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	// check if the product is the same
	require.Equal(t, product.Product.ID, product2.ID)
	require.Equal(t, product.ProductTrans.Name, product2.Name)
	require.Equal(t, product.ProductTrans.Description, product2.Description)
	require.Equal(t, product.ProductTrans.CreatedAt, product2.CreatedAt)
	require.Equal(t, product.ProductTrans.UpdatedAt, product2.UpdatedAt)
	require.Equal(t, product.ProductTrans.Category, product2.Category)
}

func TestQueries_ListProductWithInfo(t *testing.T) {
	num := 10
	langCode := LanguageCode(utils.RandomLanguage())

	// create dummy products
	for i := 0; i < num; i++ {
		addProductTxWithLangCode(t, langCode)
	}

	// list products with info and check for errors
	arg := ListProductWithInfoParams{
		Language: langCode,
		Limit:    int32(num),
		Offset:   0,
	}
	products, err := testQueries.ListProductWithInfo(context.Background(), arg)

	// check for errors or empty products
	require.NoError(t, err)
	for _, product := range products {
		require.NotEmpty(t, product)
	}

}
