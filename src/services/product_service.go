package services

import (
	"time"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func (svc *ProductService) CreateProduct(payload *dto.CreateProductDTO) (*dto.ProductDTO, error) {
	product := &models.ProductModel{
		ProductName:  payload.ProductName,
		ProductPrice: payload.ProductPrice,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	product, err := svc.productRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product.ToProductDTO(), nil
}
