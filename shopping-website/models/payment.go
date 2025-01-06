package models

import (
	"time"
)

type Payment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id"`   // 關聯的訂單 ID
	Amount    float64   `json:"amount"`     // 支付金額
	Status    string    `json:"status"`     // 支付狀態，例如 "pending", "completed", "failed"
	CreatedAt time.Time `json:"created_at"` // 支付創建時間
}
