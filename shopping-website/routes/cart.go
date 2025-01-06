package routes

import (
	"shopping-website/controllers"

	"github.com/gin-gonic/gin"
)

// SetupCartRoutes sets up the cart routes
func SetupCartRoutes(router *gin.RouterGroup, cartController *controllers.CartController) {
	router.POST("/add", cartController.AddToCart)    // 添加商品到購物車
	router.GET("/", cartController.ViewCart)         // 查看購物車
	router.PUT("/update", cartController.UpdateCart) // 更新購物車中的商品
}
