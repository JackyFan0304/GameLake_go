package routes

import (
	"shopping-website/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(r *gin.Engine, authController *controllers.AuthController, productController *controllers.ProductController, cartController *controllers.CartController, orderController *controllers.OrderController) {
	// Unregistered user routes
	authRoutes := r.Group("/auth") // 將路由組設置為 /auth
	{
		SetupAuthRoutes(authRoutes, authController) // 傳遞 authRoutes 和 authController
	}

	// Product routes
	productRoutes := r.Group("/products")
	{
		SetupProductRoutes(productRoutes, productController) // 傳遞 productController
	}

	// Cart routes
	cartRoutes := r.Group("/cart")
	{
		SetupCartRoutes(cartRoutes, cartController) // 傳遞 cartController
	}

	// Order routes
	orderRoutes := r.Group("/orders")
	{
		SetupOrderRoutes(orderRoutes, orderController) // 傳遞 cartController
	}
}
