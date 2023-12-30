package services

import (
	"time"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/dto"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/models"
	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	ProductRepository *repositories.ProductRepository
}

func (svc *ProductService) CreateProduct(payload *dto.CreateProductDTO) (*dto.ProductDTO, error) {
	product := &models.ProductModel{
		ProductName:  payload.ProductName,
		ProductPrice: payload.ProductPrice,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	product, err := svc.ProductRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product.ToProductDTO(), nil
}

func (svc *ProductService) GetProductByProductId(productId string) (*dto.ProductDTO, error) {
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, err
	}
	return svc.ProductRepository.GetProductById(objectId)
}

func (svc *ProductService) DeleteProductByProductId(productId string) error {
	objectId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil
	}
	return svc.ProductRepository.DeleteProduct(objectId)
}
