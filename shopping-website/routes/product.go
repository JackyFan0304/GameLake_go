package routes

import (
	"shopping-website/controllers"

	"github.com/gin-gonic/gin"
)

// SetupProductRoutes sets up the product routes
func SetupProductRoutes(router *gin.RouterGroup, productController *controllers.ProductController) {
	router.GET("/", productController.GetProducts) // 使用根路由
	router.GET("/:id", productController.GetProductDetail)
}
