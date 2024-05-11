package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	api := server.router.Group("/api")
	server.setupV1Routes(api)
}

func (server *Server) setupV1Routes(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	v1.GET("/healthcheck", server.healthCheck)
	server.setupV1ProductsRoutes(v1)
}

func (server *Server) setupV1ProductsRoutes(v1 *gin.RouterGroup) {

	products := v1.Group("/products")
	products.GET("", server.ListProducts)
	products.DELETE("/:id", server.DeleteProduct)
}
