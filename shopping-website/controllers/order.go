package controllers

import (
	"net/http"
	"shopping-website/models"
	"shopping-website/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService   *services.OrderService   // 使用指針類型
	PaymentService *services.PaymentService // 引入 PaymentService
}

// NewOrderController initializes a new OrderController
func NewOrderController(orderService *services.OrderService, paymentService *services.PaymentService) *OrderController {
	return &OrderController{
		OrderService:   orderService,
		PaymentService: paymentService,
	}
}

// Checkout handles the checkout process for the user
func (oc *OrderController) Checkout(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	order.UserID = uint(userID)

	// 計算總價格
	if err := order.CalculateTotalPrice(oc.OrderService.DB); err != nil { // 假設您在 OrderService 中有 DB 屬性
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total price: " + err.Error()})
		return
	}

	if err := oc.OrderService.CreateOrder(&order); err != nil { // 傳遞值而不是指針
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

// ViewOrderHistory retrieves the order history for the user
func (oc *OrderController) ViewOrderHistory(c *gin.Context) {
	userID := c.Param("userID")
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	orders, err := oc.OrderService.ViewOrderHistory(uint(userIDUint)) // 使用 ViewOrderHistory 方法
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order history: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// NotifyTransfer notifies the system of a completed transfer
func (oc *OrderController) NotifyTransfer(c *gin.Context) {
	var notification models.TransferNotification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := oc.OrderService.GetOrderByID(notification.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order: " + err.Error()})
		return
	}

	if order.Status != "Paid" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order is not paid yet"})
		return
	}

	payment := models.Payment{
		OrderID:   order.ID,
		Amount:    order.TotalPrice,
		Status:    "completed",
		CreatedAt: time.Now(),
	}

	if err := oc.PaymentService.CreatePayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record payment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer notification received"})
}
