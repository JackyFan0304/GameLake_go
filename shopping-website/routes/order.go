package routes

import (
	"shopping-website/controllers"

	"github.com/gin-gonic/gin"
)

// SetupOrderRoutes sets up the order routes
func SetupOrderRoutes(router *gin.RouterGroup, orderController *controllers.OrderController) {
	router.POST("/checkout", orderController.Checkout)              // 處理結帳
	router.GET("/history", orderController.ViewOrderHistory)        // 獲取訂單歷史
	router.POST("/notify-transfer", orderController.NotifyTransfer) // 通知轉帳
}
