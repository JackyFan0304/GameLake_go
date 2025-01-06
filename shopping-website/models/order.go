package models

import (
	"time"

	"gorm.io/gorm"
)

// OrderItem 模型定義
type OrderItem struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	OrderID   uint `json:"order_id"`   // 關聯的訂單 ID
	ProductID uint `json:"product_id"` // 關聯的產品 ID
	Quantity  int  `json:"quantity"`   // 訂購數量
}

// Order 模型定義
type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserID     uint        `json:"user_id"`
	TotalPrice float64     `json:"total_price"`                      // 總價格
	Status     string      `json:"status"`                           // e.g., "pending", "completed", "canceled"
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"` // 使用 time.Time 並自動設置
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"` // 使用 time.Time 並自動更新
	OrderItems []OrderItem `json:"order_items"`                      // 關聯的產品項目
}

// CalculateTotalPrice calculates the total price of the order based on order items
func (o *Order) CalculateTotalPrice(db *gorm.DB) error {
	var total float64
	for _, item := range o.OrderItems {
		productPrice, err := GetProductPrice(item.ProductID, db) // 獲取產品價格
		if err != nil {
			return err // 返回錯誤
		}
		total += productPrice * float64(item.Quantity) // 計算總價格
	}
	o.TotalPrice = total // 設置訂單總價格
	return nil
}

// GetProductPrice retrieves the price of a product by its ID (應該在 services 包中)
func GetProductPrice(productID uint, db *gorm.DB) (float64, error) {
	var product Product // 假設您有一個 Product 模型，並且它包含 Price 字段
	if err := db.First(&product, productID).Error; err != nil {
		return 0, err // 如果找不到產品，返回錯誤
	}
	return product.Price, nil // 返回產品價格
}
