package models

import "time"

type TransferNotification struct {
	OrderID   uint      `json:"order_id"`   // 訂單 ID
	Status    string    `json:"status"`     // 轉帳狀態，例如 "completed", "pending"
	CreatedAt time.Time `json:"created_at"` // 通知創建時間
}
