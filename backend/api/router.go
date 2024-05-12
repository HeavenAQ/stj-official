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
	server.setupV1UsersRoutes(v1)
	server.setupV1AuthRoutes(v1)
}

func (server *Server) setupV1ProductsRoutes(v1 *gin.RouterGroup) {

	products := v1.Group("/products")
	products.GET("", server.ListProducts)
	products.DELETE("/:id", server.DeleteProduct)
}

func (server *Server) setupV1UsersRoutes(v1 *gin.RouterGroup) {
	//product := v1.Group("/users")
	//product.GET("/:id", server.GetUser)
	//product.DELETE("/:id", server.DeleteUser)
	//product.PUT("/:id", server.UpdateUser)
}

func (server *Server) setupV1AuthRoutes(v1 *gin.RouterGroup) {
	auth := v1.Group("/auth")
	auth.POST("/register", server.UserRegister)
	auth.POST("/login", server.UserLogin)
	//auth.GET("/logout", server.Logout)
	//auth.POST("/password-reset", server.PasswordReset)
}
