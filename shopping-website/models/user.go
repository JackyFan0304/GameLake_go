package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}

// UserModel 提供用戶數據庫操作的方法
type UserModel struct {
	DB *gorm.DB
}

// NewUserModel 創建新的 UserModel 實例
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{DB: db}
}

// CreateUser 創建新用戶
func (m *UserModel) CreateUser(user *User) error {
	result := m.DB.Create(user)
	return result.Error
}

// FindUser 根據電子郵件查找用戶
func (m *UserModel) FindUser(email string) (*User, error) {
	var user User
	result := m.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 用戶不存在時返回 nil
		}
		return nil, result.Error // 返回其他錯誤
	}
	return &user, nil // 找到用戶時返回用戶
}
