package controllers

import (
	"net/http"
	"shopping-website/services"
	"strconv" // 引入 strconv 以便進行字符串轉整數的轉換

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *services.ProductService // 使用指針類型
}

// GetProducts retrieves all products
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductDetail retrieves a product by its ID
func (pc *ProductController) GetProductDetail(c *gin.Context) {
	idStr := c.Param("id")                      // 獲取字符串形式的 ID
	id, err := strconv.ParseUint(idStr, 10, 32) // 將字符串轉換為 uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := pc.ProductService.GetProductByID(uint(id)) // 使用 uint(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
