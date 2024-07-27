package api

import (
	"fmt"
	"net/http"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/helpers"
	"stj-ecommerce/token"

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
	user, err := helpers.AuthAndGetUser(ctx, req, server.store, authorizationHeaderKey)
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
	user, err := helpers.AuthAndGetUser(ctx, req, server.store, authorizationHeaderKey)
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
		product, err := server.store.GetProductByID(ctx, item.ProductID)
		if err != nil {
			server.ErrorLogger.Println(err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("product %v not found", item.Name)})
			return
		}

		// create order detail
		arg := db.CreateOrderDetailsParams{
			OrderPk:   order.Pk,
			ProductPk: product.Pk,
			Quantity:  item.Quantity,
			Price:     0,
		}
		_, err = server.store.CreateOrderDetails(ctx, arg)
		if err != nil {
			server.ErrorLogger.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create order detail for product %v", item.Name)})
			return
		}

		// add to order item
		items[i] = helpers.OrderItem{
			ProductID: product.ID,
			Name:      item.Name,
			Quantity:  item.Quantity,
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

type GetOrderRequest struct {
	OrderID pgtype.UUID `json:"order_id" binding:"required"`
}

func (server *Server) GetOrder(ctx *gin.Context) {
	// get order id from uri
	var req GetOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err)
		return
	}

	// get user
	userID := ctx.MustGet(authorizationHeaderKey).(*token.Payload).QueryID
	user, err := server.store.GetUserByID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to get user info: %v", err.Error())})
		server.ErrorLogger.Println(err.Error())
		return
	}

	// get order by user and id
	args := db.GetOrderByUserAndOrderIDParams{
		UserPk: user.Pk,
		ID:     req.OrderID,
	}
	order, err := server.store.GetOrderByUserAndOrderID(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err.Error())
		return
	}

	// get order items
	orderItems, err := helpers.GetOrderItemsFromDB(ctx, server.store, order.Pk, user.Language)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		server.ErrorLogger.Println(err.Error())
		return
	}

	// send response
	ctx.JSON(http.StatusOK, GetOrderResponse{
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		Items:     orderItems,
	})
}
