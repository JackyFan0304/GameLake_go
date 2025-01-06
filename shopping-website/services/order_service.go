package services

import (
	"errors"
	"shopping-website/models"

	"gorm.io/gorm"
)

type OrderService struct {
	DB *gorm.DB // 直接在 OrderService 中使用 DB
}

// NewOrderService initializes a new OrderService
func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		DB: db,
	}
}

// CreateOrder creates a new order in the database
func (os *OrderService) CreateOrder(order *models.Order) error {
	if err := os.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// GetOrderByID retrieves an order by its ID
func (os *OrderService) GetOrderByID(orderID uint) (models.Order, error) {
	var order models.Order
	if err := os.DB.First(&order, orderID).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

// GetOrdersByUserID retrieves all orders for a specific user
func (os *OrderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := os.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// ViewOrderHistory retrieves all orders for a specific user
func (os *OrderService) ViewOrderHistory(userID uint) ([]models.Order, error) {
	return os.GetOrdersByUserID(userID)
}

// NotifyTransfer checks the order status and performs notification logic
func (os *OrderService) NotifyTransfer(orderID uint) error {
	order, err := os.GetOrderByID(orderID)
	if err != nil {
		return err
	}
	if order.Status != "Paid" {
		return errors.New("order is not paid yet")
	}

	// Logic to notify transfer (e.g., send email or update status)
	// 這裡可以根據需求實現具體的通知邏輯

	return nil
}
