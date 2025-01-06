package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

// ProductModel 用於處理與產品相關的數據庫操作
type ProductModel struct {
	db *gorm.DB // 資料庫連接
}

// NewProductModel 構造函數，用於初始化 ProductModel
func NewProductModel(db *gorm.DB) *ProductModel {
	return &ProductModel{db: db}
}

// GetAllProducts 獲取所有產品
func (pm *ProductModel) GetAllProducts() ([]Product, error) {
	var products []Product
	if err := pm.db.Find(&products).Error; err != nil {
		return nil, err // 返回 nil 和錯誤
	}
	return products, nil // 返回產品列表和 nil 錯誤
}

// GetProductByID 根據 ID 獲取產品
func (pm *ProductModel) GetProductByID(id uint) (Product, error) {
	var product Product
	if err := pm.db.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return Product{}, fmt.Errorf("product not found") // 返回空的 Product 和錯誤
		}
		return Product{}, err // 返回空的 Product 和其他錯誤
	}
	return product, nil // 返回找到的產品和 nil 錯誤
}
