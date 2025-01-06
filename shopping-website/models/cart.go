package models

import "errors"

type Cart struct {
	ID     string     `json:"id"`
	UserID string     `json:"user_id"`
	Items  []CartItem `json:"items"`
}

// CartItem represents an item in the cart
type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(productID string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	// Check if the item already exists in the cart
	for i, item := range c.Items {
		if item.ProductID == productID {
			c.Items[i].Quantity += quantity // Update quantity
			return nil
		}
	}

	// If it doesn't exist, add a new CartItem
	c.Items = append(c.Items, CartItem{ProductID: productID, Quantity: quantity})
	return nil
}

// GetCartItems retrieves items in the cart
func (c *Cart) GetCartItems() []CartItem {
	return c.Items // Return the list of items in the cart
}
