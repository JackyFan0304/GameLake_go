package services

import (
	"errors"
	"shopping-website/models"
)

type ProductService struct {
	productModel *models.ProductModel // 使用 models.ProductModel
}

// NewProductService 構造函數，用於初始化 ProductService
func NewProductService(productModel *models.ProductModel) *ProductService {
	return &ProductService{
		productModel: productModel,
	}
}

func (ps *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := ps.productModel.GetAllProducts() // 調用 ProductModel 的方法
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProductByID(id uint) (models.Product, error) { // 將 id 類型改為 uint
	product, err := ps.productModel.GetProductByID(id) // 調用 ProductModel 的方法
	if err != nil {
		return models.Product{}, err
	}
	if (product == models.Product{}) { // 檢查產品是否存在
		return models.Product{}, errors.New("product not found")
	}
	return product, nil
}
