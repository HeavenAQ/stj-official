package db

import (
	"context"
	"stj-ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func addRandomProductTx(t *testing.T) *ProductTxResult {
	arg := AddProductTxParams{
		LangCode:    LanguageCode(utils.RandomLanguage()),
		Name:        utils.RandomAlphabetString(10),
		Description: utils.RandomAlphabetString(20),
		ImageURLs:   []string{utils.RandomAlphabetString(10), utils.RandomAlphabetString(10)},
		Price:       int32(utils.RandomInt(0, 1000)),
		Status:      ProductStatusInStock,
		Quantity:    10,
		Category:    "test",
	}
	return addProductTx(t, arg)
}

func addProductTxWithLangCode(t *testing.T, langCode LanguageCode) *ProductTxResult {
	arg := AddProductTxParams{
		LangCode:    langCode,
		Name:        utils.RandomAlphabetString(10),
		Description: utils.RandomAlphabetString(20),
		ImageURLs:   []string{utils.RandomAlphabetString(10), utils.RandomAlphabetString(10)},
		Price:       int32(utils.RandomInt(0, 1000)),
		Status:      ProductStatusInStock,
		Quantity:    10,
		Category:    "test",
	}
	return addProductTx(t, arg)
}

func addProductTx(t *testing.T, arg AddProductTxParams) *ProductTxResult {
	// ensure the product is added
	product, err := testStore.AddProductTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	// ensure the product description is correct
	require.Equal(t, arg.Name, product.ProductTrans.Name)
	require.Equal(t, arg.Description, product.ProductTrans.Description)
	require.Equal(t, arg.LangCode, product.ProductTrans.Language)
	require.Equal(t, arg.Category, product.ProductTrans.Category)
	require.NotZero(t, product.ProductTrans.Pk)
	require.NotZero(t, product.ProductTrans.CreatedAt)
	require.NotZero(t, product.ProductTrans.UpdatedAt)

	// ensure the product is correct
	require.Equal(t, arg.Quantity, product.Product.Quantity)
	require.Equal(t, arg.Status, product.Product.Status)
	require.Equal(t, arg.Price, product.Product.Price)
	require.Equal(t, arg.ImageURLs, product.Product.ImageURLs)
	require.NotZero(t, product.Product.ID)
	require.NotZero(t, product.Product.Pk)
	require.NotZero(t, product.Product.CreatedAt)
	require.NotZero(t, product.Product.UpdatedAt)
	return product
}

func TestStore_NewStore(t *testing.T) {
	require.NotNil(t, testStore)
	require.NotNil(t, testStore.Queries)
	require.NotNil(t, testStore.db)
}

func TestStore_HealthCheck(t *testing.T) {
	err := testStore.HealthCheck()
	require.NoError(t, err)
}

func TestStore_AddProductTx(t *testing.T) {
	product := addRandomProductTx(t)
	testQueries.DeleteProduct(context.Background(), product.Product.Pk)
}
