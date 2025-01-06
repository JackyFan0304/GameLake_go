package services

import (
	"shopping-website/models"

	"gorm.io/gorm"
)

type PaymentService struct {
	DB *gorm.DB
}

// CreatePayment creates a new payment record in the database
func (ps *PaymentService) CreatePayment(payment models.Payment) error {
	if err := ps.DB.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

// GetPaymentByOrderID retrieves the payment record for a specific order
func (ps *PaymentService) GetPaymentByOrderID(orderID uint) (models.Payment, error) {
	var payment models.Payment
	if err := ps.DB.Where("order_id = ?", orderID).First(&payment).Error; err != nil {
		return models.Payment{}, err
	}
	return payment, nil
}
