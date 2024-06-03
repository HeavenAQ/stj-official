package helpers

import (
	db "stj-ecommerce/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderItem struct {
	Name     string      `json:"name"`
	Quantity int32       `json:"quantity"`
	ID       pgtype.UUID `json:"product_id"`
}

func GetOrderItemsFromDB(ctx *gin.Context, store *db.Store, orderID int64, language db.LanguageCode) ([]OrderItem, error) {
	// get order details
	orderDetails, err := store.GetOrderDetailByOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}

	// get items
	allItems := make([]OrderItem, len(orderDetails))
	for j, detail := range orderDetails {
		// get product
		product, err := store.GetProduct(ctx, detail.ProductPk)
		if err != nil {
			return nil, err
		}

		// get product information
		arg := db.GetProductTranslationParams{
			ProductPk: product.Pk,
			Language:  language,
		}
		productInfo, err := store.GetProductTranslation(ctx, arg)
		if err != nil {
			return nil, err
		}

		// create item
		allItems[j] = OrderItem{
			ID:       product.ID,
			Name:     productInfo.Name,
			Quantity: detail.Quantity,
		}
	}

	return allItems, nil
}
