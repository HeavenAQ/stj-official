package api

import (
	"fmt"
	"net/http"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/helpers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type ListOrdersRequest struct {
	Limit  int32 `json:"limit" binding:"required"`
	Offset int32 `json:"offset" binding:"required"`
}

type GetOrderResponse struct {
	Status    db.OrderStatus      `json:"status"`
	CreatedAt pgtype.Timestamptz  `json:"created_at"`
	Items     []helpers.OrderItem `json:"items"`
}

func (server *Server) ListOrders(ctx *gin.Context) {
	var req ListOrdersRequest
	user, err := helpers.VerifyJSONAndGetUser(ctx, req, server.store, authorizationHeaderKey)
	if err != nil {
		server.ErrorLogger.Println(err)
		return
	}

	// get a user's orders
	arg := db.ListOrdersByUserParams{
		UserPk: user.Pk,
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	orders, err := server.store.ListOrdersByUser(ctx, arg)
	if err != nil {
		server.ErrorLogger.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list orders"})
		return
	}

	// get orders' details
	listOrderResponse := make([]GetOrderResponse, len(orders))
	for i, order := range orders {
		// get items
		items, err := helpers.GetOrderItemsFromDB(ctx, server.store, order.Pk, user.Language)
		if err != nil {
			server.ErrorLogger.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get items for order %v", order.ID)})
			return
		}

		// add to order response
		listOrderResponse[i] = GetOrderResponse{
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
			Items:     items,
		}
	}

	// return orders
	ctx.JSON(http.StatusOK, listOrderResponse)
}

type CreateOrderRequest struct {
	ShippingAddress string              `json:"shipping_address" binding:"required"`
	Phone           string              `json:"phone" binding:"required"`
	Email           string              `json:"email" binding:"required"`
	Items           []helpers.OrderItem `json:"items" binding:"required"`
}

type CreateOrderResponse struct {
	CreatedAt   pgtype.Timestamptz  `json:"created_at"`
	OrderStatus db.OrderStatus      `json:"status"`
	Items       []helpers.OrderItem `json:"items"`
	OrderID     pgtype.UUID         `json:"order_id"`
}

func (server *Server) CreateOrder(ctx *gin.Context) {
	var req CreateOrderRequest
	user, err := helpers.VerifyJSONAndGetUser(ctx, req, server.store, authorizationHeaderKey)
	if err != nil {
		server.ErrorLogger.Println(err)
		return
	}

	// create order
	arg := db.CreateOrderParams{
		UserPk:          user.Pk,
		Status:          db.OrderStatusPending,
		TotalPrice:      0,
		ShippingAddress: req.ShippingAddress,
		ShippingDate:    pgtype.Timestamptz{},
		DeliveredDate:   pgtype.Timestamptz{},
		Phone:           req.Phone,
		Email:           req.Email,
		IsPaid:          false,
	}
	order, err := server.store.CreateOrder(ctx, arg)
	if err != nil {
		server.ErrorLogger.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	// create all order details (link products to order)
	items := make([]helpers.OrderItem, len(req.Items))
	for i, item := range req.Items {
		// ensure the product exists
		product, err := server.store.GetProductByID(ctx, item.ID)
		if err != nil {
			server.ErrorLogger.Println(err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("product %v not found", item.Name)})
			return
		}

		// create order detail
		arg := db.CreateOrderDetailParams{
			OrderPk:   order.Pk,
			ProductPk: product.Pk,
			Quantity:  item.Quantity,
			Price:     0,
		}
		_, err = server.store.CreateOrderDetail(ctx, arg)
		if err != nil {
			server.ErrorLogger.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create order detail for product %v", item.Name)})
			return
		}

		// add to order item
		items[i] = helpers.OrderItem{
			ID:       product.ID,
			Name:     item.Name,
			Quantity: item.Quantity,
		}
	}

	// return order response
	ctx.JSON(http.StatusOK, CreateOrderResponse{
		OrderID:     order.ID,
		OrderStatus: order.Status,
		CreatedAt:   order.CreatedAt,
		Items:       items,
	})
}
