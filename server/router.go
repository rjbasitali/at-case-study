package server

import (
	"git.rjbasitali.com/at-case-study/cfg"
	"git.rjbasitali.com/at-case-study/controllers/auth/token"
	"git.rjbasitali.com/at-case-study/controllers/products"
	"git.rjbasitali.com/at-case-study/middlewares"
	"github.com/gin-gonic/gin"
)

// newRouter creates a new router.
// It registers the following routes and middleware:
// - GET /auth/token
// - AuthMiddleware
// - GET /products/:id
// - POST /products
func newRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/auth/token", token.GetToken)
	r.Use(middlewares.AuthMiddleware(cfg.JWT_SIGNING_SERCRET))
	r.GET("/products/:id", products.GetProduct)
	r.POST("/products", products.CreateProduct)

	return r
}
