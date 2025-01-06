package services

import (
	"errors"
	"shopping-website/models"
	"sync"
)

type CartService struct {
	carts map[string]*models.Cart
	mu    sync.Mutex
}

func NewCartService() *CartService {
	return &CartService{
		carts: make(map[string]*models.Cart),
	}
}

// AddItem adds an item to the shopping cart
func (cs *CartService) AddItem(userID string, productID string, quantity int) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	cart, exists := cs.carts[userID]
	if !exists {
		cart = &models.Cart{ID: userID, UserID: userID, Items: []models.CartItem{}}
		cs.carts[userID] = cart
	}

	// 檢查產品是否已經在購物車中
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity += quantity // 增加數量
			return nil
		}
	}

	// 如果不存在，添加新的 CartItem
	cart.Items = append(cart.Items, models.CartItem{ProductID: productID, Quantity: quantity})
	return nil
}

// ViewCart retrieves the cart for a specific user
func (cs *CartService) ViewCart(userID string) (*models.Cart, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cart, exists := cs.carts[userID]
	if !exists {
		return nil, errors.New("cart not found")
	}

	return cart, nil
}

// UpdateCart updates the items in the shopping cart
func (cs *CartService) UpdateCart(userID string, productID string, quantity int) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cart, exists := cs.carts[userID]
	if !exists {
		return errors.New("cart not found")
	}

	if quantity <= 0 {
		// 如果數量為0，則刪除該項目
		for i := 0; i < len(cart.Items); i++ {
			if cart.Items[i].ProductID == productID {
				cart.Items = append(cart.Items[:i], cart.Items[i+1:]...) // 刪除項目
				break
			}
		}
	} else {
		for i := range cart.Items {
			if cart.Items[i].ProductID == productID {
				cart.Items[i].Quantity = quantity // 更新數量
				return nil
			}
		}

		// 如果不存在，添加新的 CartItem
		cart.Items = append(cart.Items, models.CartItem{ProductID: productID, Quantity: quantity})
	}

	return nil
}
