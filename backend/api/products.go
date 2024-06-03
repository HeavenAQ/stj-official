package api

import (
	"net/http"
	db "stj-ecommerce/db/sqlc"

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

type CreateProductRequest struct {
	JPName         string
	ChnName        string
	JpCategory     string
	ChnCategory    string
	StockStatus    db.ProductStatus
	jpDescription  string
	chnDescription string
	Price          int64
	quantity       int64
}
