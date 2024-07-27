package api

import (
	"net/http"
	db "stj-ecommerce/db/sqlc"
	"stj-ecommerce/helpers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type ListProductsRequest struct {
	Language string `form:"language" binding:"required"`
	Limit    int32  `form:"limit" binding:"required"`
	Offset   int32  `form:"offset"`
}

func (server *Server) ListProducts(ctx *gin.Context) {
	var req ListProductsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create args for listing products
	args := db.ListProductWithInfoParams{
		Language: db.LanguageCode(req.Language),
		Limit:    req.Limit,
		Offset:   req.Offset,
	}

	// list products from database
	products, err := server.store.ListProductWithInfo(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

type DeleteProductRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) DeleteProduct(ctx *gin.Context) {
	var req DeleteProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// convert string to uuid
	var pgUUID pgtype.UUID
	pgUUID.Scan(req.ID)

	// delete product from database
	err := server.store.DeleteProductById(ctx, pgUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

// create product
type Description struct {
	Introduction   string `json:"introduction" binding:"required"`
	Prize          string `json:"prize" binding:"required"`
	ItemInfo       string `json:"item_info" binding:"required"`
	Recommendation string `json:"recommendation" binding:"required"`
}

type CreateProductRequest struct {
	Name        string           `json:"name" binding:"required"`
	Language    string           `json:"language" binding:"required"`
	Category    string           `json:"category" binding:"required"`
	Status      db.ProductStatus `json:"status"`
	Description Description      `json:"description" binding:"required"`
	ImageURLs   []string         `json:"image_urls" binding:"required"`
	Price       int32            `json:"price" binding:"required"`
	Quantity    int32            `json:"quantity" binding:"required"`
	IsHot       bool             `json:"is_hot"`
}

type CreateProductResponse struct {
	Name        string           `json:"name"`
	Language    string           `json:"language"`
	Category    string           `json:"category"`
	Status      db.ProductStatus `json:"status"`
	Description Description      `json:"description"`
	ImageURLs   []string         `json:"image_urls"`
	Price       int32            `json:"price"`
	Quantity    int32            `json:"quantity"`
	IsHot       bool             `json:"is_hot"`
}

func (server *Server) CreateProduct(ctx *gin.Context) {
	// ensure the incoming request is valid and get the user
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.AuthAndGetUser(ctx, req, server.store, authorizationPayloadKey)
	if err != nil {
		server.ErrorLogger.Println(err)
		return
	}

	// check if the user is an admin
	if user.Email != server.config.ADMIN_EMAIL {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		server.WarnLogger.Println("Attempt to create product by non-admin")
		return
	}

	// start create-product transaction
	product, err := server.store.CreateProductTx(ctx, db.CreateProductTxParams{
		LangCode:       db.LanguageCode(req.Language),
		Name:           req.Name,
		Status:         req.Status,
		Category:       req.Category,
		ImageURLs:      req.ImageURLs,
		Price:          req.Price,
		Quantity:       req.Quantity,
		Introduction:   req.Description.Introduction,
		Prize:          req.Description.Prize,
		ItemInfo:       req.Description.ItemInfo,
		Recommendation: req.Description.Recommendation,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return the result back to user
	ctx.JSON(http.StatusOK, CreateProductResponse{
		Name:     product.ProductTrans.Name,
		Language: string(product.ProductTrans.Language),
		Category: product.ProductTrans.Category,
		Status:   product.Product.Status,
		Description: Description{
			Introduction:   product.ProductDesc.Introduction,
			Prize:          product.ProductDesc.Prize,
			ItemInfo:       product.ProductDesc.ItemInfo,
			Recommendation: product.ProductDesc.Recommendation,
		},
		ImageURLs: product.Product.ImageURLs,
		Price:     product.Product.Price,
		Quantity:  product.Product.Quantity,
		IsHot:     product.Product.IsHot,
	})
}
