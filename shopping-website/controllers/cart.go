package controllers

import (
	"net/http"
	"shopping-website/models"
	"shopping-website/services"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	CartService *services.CartService // 使用指針類型
}

// AddToCart adds an item to the shopping cart
func (cc *CartController) AddToCart(c *gin.Context) {
	var item models.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("user_id") // 假設您從請求參數中獲取 user_id

	if err := cc.CartService.AddItem(userID, item.ProductID, item.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

// ViewCart retrieves the items in the shopping cart
func (cc *CartController) ViewCart(c *gin.Context) {
	userID := c.Param("user_id") // 假設您從請求參數中獲取 user_id
	cartItems, err := cc.CartService.ViewCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartItems)
}

// UpdateCart updates the items in the shopping cart
func (cc *CartController) UpdateCart(c *gin.Context) {
	var items []models.CartItem
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("user_id") // 假設您從請求參數中獲取 user_id

	for _, item := range items {
		if err := cc.CartService.UpdateCart(userID, item.ProductID, item.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart updated"})
}
