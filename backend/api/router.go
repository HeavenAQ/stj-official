package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	api := server.router.Group("/api")
	server.setupV1Routes(api)
}

func (server *Server) setupV1Routes(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	v1.GET("/healthcheck", server.healthCheck)
	server.setupV1ProductsRoutes(v1)
	server.setupV1UsersRoutes(v1)
	server.setupV1AuthRoutes(v1)
}

func (server *Server) setupV1ProductsRoutes(v1 *gin.RouterGroup) {
	products := v1.Group("/products")
	products.GET("", server.ListProducts)
	products.GET("/:id", server.GetProduct)

	// Protected routes
	products.Use(server.authMiddleware(server.tokenMaker))
	products.POST("", server.CreateProduct)
	products.DELETE("/:id", server.DeleteProduct)
}

func (server *Server) setupV1UsersRoutes(v1 *gin.RouterGroup) {
	users := v1.Group("/users").Use(server.authMiddleware(server.tokenMaker))
	users.GET("", server.GetUser)
	users.PUT("", server.UpdateUser)
	// product.DELETE("/:id", server.DeleteUser)
}

func (server *Server) setupV1OrdersRoutes(v1 *gin.RouterGroup) {
	orders := v1.Group("/orders").Use(server.authMiddleware(server.tokenMaker))
	orders.GET("", server.ListOrders)
	orders.GET("/:id", server.GetOrder)
	orders.POST("", server.CreateOrder)
}

func (server *Server) setupV1AuthRoutes(v1 *gin.RouterGroup) {
	auth := v1.Group("/auth")
	auth.POST("/register", server.UserRegister)
	auth.POST("/login", server.UserLogin)
	auth.POST("/refresh", server.RefreshAccess)
	// auth.GET("/logout", server.Logout)
	// auth.POST("/password-reset", server.PasswordReset)
}
